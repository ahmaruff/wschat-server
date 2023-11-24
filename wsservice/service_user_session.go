package wsservice

import (
	"errors"

	"github.com/oklog/ulid/v2"
)

var UserSessionList = make(map[string]*UserSession)
var ErrUserSessionNotFound = errors.New("user session not found")

func FindUserSession(id ulid.ULID) (*UserSession, error) {
	userSession, ok := UserSessionList[id.String()]
	if !ok {
		return nil, ErrUserSessionNotFound
	}
	return userSession, nil
}

func FindUserSessionByUserId(id ulid.ULID) (*UserSession, error) {
	for _, uS := range UserSessionList {
		if uS.User.Id == id {
			return uS, nil
		}
	}
	return nil, ErrUserSessionNotFound

}

func AddToUserSessionList(userSession *UserSession) {
	// mutex.Lock()
	// defer mutex.Unlock()

	UserSessionList[userSession.Id.String()] = userSession
}

func RemoveUserSession(userSessionId ulid.ULID) {
	// mutex.Lock()
	// defer mutex.Unlock()

	delete(UserSessionList, userSessionId.String())
}
