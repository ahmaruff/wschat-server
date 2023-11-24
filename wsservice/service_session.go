package wsservice

import (
	"errors"

	"github.com/oklog/ulid/v2"
)

var SessionList = make(map[string]*Session)
var ErrSessionNotFound = errors.New("session not found")

func FindSession(id ulid.ULID) (*Session, error) {
	session, ok := SessionList[id.String()]
	if !ok {
		return nil, ErrSessionNotFound
	}
	return session, nil
}

func AddToSessionList(session *Session) {
	// mutex.Lock()
	// defer mutex.Unlock()

	SessionList[session.Id.String()] = session
}

func RemoveSession(sessionId ulid.ULID) {
	// mutex.Lock()
	// defer mutex.Unlock()

	delete(SessionList, sessionId.String())
}
