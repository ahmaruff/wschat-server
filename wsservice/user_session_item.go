package wsservice

import (
	"ahmaruff/wschat/user"

	"github.com/gorilla/websocket"
	"github.com/oklog/ulid/v2"
)

type UserSession struct {
	Id ulid.ULID
	*user.User
	*websocket.Conn
}

func MakeNewUserSession(u *user.User, ws *websocket.Conn) UserSession {
	s := UserSession{
		Id:   ulid.Make(),
		User: u,
		Conn: ws,
	}

	return s
}
