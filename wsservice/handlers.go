package wsservice

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"

	gubrak "github.com/novalagung/gubrak/v2"
	"github.com/rs/zerolog/log"
)

type M map[string]interface{}

const MSG_NEW_USER = "New User"
const MSG_CHAT = "Chat"
const MSG_LEAVE = "Leave"

type WebsocketConnection struct {
	*websocket.Conn
	Username string
}

var connections = make([]*WebsocketConnection, 0)

type SocketPayload struct {
	Message string `json:"message"`
}

type SocketResponse struct {
	From    string `json:"sender"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{}

func WebsocketHandler(c echo.Context) error {
	wsConn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Error().Err(err).Msg("ws: Could not open the websocket connection!")
		return echo.NewHTTPError(http.StatusInternalServerError, "ws: Could not open the websocket connection!")
	}

	username := c.QueryParam("username")

	currentConn := WebsocketConnection{
		Conn:     wsConn,
		Username: username,
	}
	connections = append(connections, &currentConn)

	go handleIo(&currentConn, connections)

	return nil
}

func ejectConnection(currentConn *WebsocketConnection) {
	filtered := gubrak.From(connections).Reject(func(each *WebsocketConnection) bool {
		return each == currentConn
	}).Result()
	connections = filtered.([]*WebsocketConnection)
}

func broadcastMessage(currentConn *WebsocketConnection, kind, message string) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}
		eachConn.WriteJSON(SocketResponse{
			From:    currentConn.Username,
			Type:    kind,
			Message: message,
		})
	}
}

func handleIo(currentConn *WebsocketConnection, connections []*WebsocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msg(fmt.Sprintf("%v", r))
		}
	}()

	broadcastMessage(currentConn, MSG_NEW_USER, "")

	for {
		payload := SocketPayload{}

		err := currentConn.ReadJSON(&payload)

		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(currentConn, MSG_LEAVE, "")
				ejectConnection(currentConn)
			}
			log.Error().Err(err)
			continue
		}
		broadcastMessage(currentConn, MSG_CHAT, payload.Message)
	}
}
