package wsservice

import (
	"ahmaruff/wschat/user"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/oklog/ulid/v2"

	"github.com/rs/zerolog/log"
)

type M map[string]interface{}

const MSG_NEW_USER = "New User"
const MSG_CHAT = "Chat"
const MSG_LEAVE = "Leave"

// var mutex = &sync.Mutex{}

type SocketPayload struct {
	SenderId string `json:"sender_id"`
	Message  string `json:"message"`
}

type SocketResponse struct {
	Sender  string `json:"sender"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func joinSession(sessionId ulid.ULID, userId ulid.ULID) {
	// mutex.Lock()
	// defer mutex.Unlock()

	if session, ok := SessionList[sessionId.String()]; ok {
		u, err := user.FindUser(userId)
		if err != nil {
			log.Error().Msg("session: user <" + userId.String() + "> not found")
		}
		session.Participants[userId] = u
		session.LastActive = time.Now()
	} else {
		log.Error().Msg("session: session <" + sessionId.String() + "> not found")
	}
}

func broadcastMessage(sessionId ulid.ULID, senderId ulid.ULID, kind, message string) {
	// mutex.Lock()
	// defer mutex.Unlock()

	if session, ok := SessionList[sessionId.String()]; ok {
		senderSession, err := FindUserSessionByUserId(senderId)
		if err != nil {
			log.Error().Msg("session: session not found")
		}

		for _, participant := range session.Participants {
			if participant.Id == senderSession.User.Id {
				continue
			}
			parSes, err := FindUserSessionByUserId(participant.Id)
			if err != nil {
				log.Error().Msg("session: session  not found")
			}
			parSes.Conn.WriteJSON(SocketResponse{
				Sender:  senderSession.Name,
				Type:    kind,
				Message: message,
			})
		}
		session.LastActive = time.Now()
	} else {
		log.Error().Msg("session: session <" + sessionId.String() + "> not found")
	}
}

var upgrader = websocket.Upgrader{}

func handleIo(currentSession *Session, currentUser *user.User, conn *websocket.Conn) {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msg(fmt.Sprintf("%v", r))
		}
	}()

	broadcastMessage(currentSession.Id, currentUser.Id, MSG_NEW_USER, "New user join the room chat")

	for {
		payload := SocketPayload{}

		err := conn.ReadJSON(&payload)

		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				s, _ := FindUserSessionByUserId(currentUser.Id)
				s.Conn.Close()
				broadcastMessage(currentSession.Id, currentUser.Id, MSG_LEAVE, "")
				RemoveSession(currentSession.Id)
			}
			log.Error().Err(err)
			continue
		}
		senderId, err := ulid.Parse(payload.SenderId)

		if err != nil {
			log.Error().Err(err)
		}

		broadcastMessage(currentSession.Id, senderId, MSG_CHAT, payload.Message)
	}
}
