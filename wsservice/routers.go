package wsservice

import (
	"ahmaruff/wschat/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func InitWsRoutes(e *echo.Echo) {
	e.GET("/session", ListSessionHandler)
	e.GET("/session/create", CreateNewSessionHandler)
	e.GET("/ws", WebsocketHandler)
}

func ListSessionHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, SessionList)
}

func CreateNewSessionHandler(c echo.Context) error {
	t := c.QueryParam("type")
	sessionId := ulid.Make()

	session, err := MakeNewSession(sessionId, t)
	if err != nil {
		return err
	}

	AddToSessionList(&session)
	return c.JSON(http.StatusOK, session)
}

func WebsocketHandler(c echo.Context) error {
	// CHECK REQUIRED PARAMETER FIRST
	// CHECK SESSION_ID
	sessionIdQuery := c.QueryParam("session_id")
	if sessionIdQuery == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "session_id empty")
	}

	// CHECK USER_ID
	userIdQuery := c.QueryParam("user_id")
	if userIdQuery == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "user_id empty")
	}

	// PARSE ULID SESSION_ID
	currentSessionID, err := ulid.Parse(sessionIdQuery)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid session_id")
	}
	currentSession, err := FindSession(currentSessionID)
	if err != nil {
		b, _ := currentSession.MarshalJSON()
		log.Debug().Bytes("current_user", b)
		return echo.NewHTTPError(http.StatusBadRequest, "session not found")
	}

	// PARSE ULID USER_ID
	userId, err := ulid.Parse(userIdQuery)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user_id")
	}

	currentUser, err := user.FindUser(userId)
	if err != nil {
		b, _ := currentUser.MarshalJSON()
		log.Debug().Bytes("current_user", b)
		return echo.NewHTTPError(http.StatusBadRequest, "user not found")
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	wsConn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Error().Err(err).Msg("ws: Could not open the websocket connection!")
		// return err
		// defer wsConn.Close()
		return echo.NewHTTPError(http.StatusInternalServerError, "ws: Could not open the websocket connection!")
	}

	joinSession(currentSession.Id, currentUser.Id)

	currentUserSession := MakeNewUserSession(currentUser, wsConn)
	currentSession.Participants[currentUser.Id] = currentUser

	AddToUserSessionList(&currentUserSession)
	// UserSessionList[currentUserSession.Id.String()] = &currentUserSession

	handleIo(currentSession, currentUser, wsConn)

	return nil
}
