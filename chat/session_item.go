package chat

import (
	"errors"
	"strings"

	"github.com/oklog/ulid/v2"
)

const PrivateSession = "PRIVATE"
const GroupSession = "GROUP"

var ErrInvalidSessionType = errors.New("chat: invalid session type")

type Session struct {
	Id   ulid.ULID
	Type string
}

func MakeNewSession(t string) (Session, error) {
	new_t, err := validateType(t)

	if err != nil {
		return Session{}, err
	}

	s := Session{
		Id:   ulid.Make(),
		Type: new_t,
	}

	return s, nil

}

func validateType(t string) (string, error) {
	var res string
	t = strings.ToUpper(t)

	switch {
	case t == PrivateSession:
		res = PrivateSession
	case t == GroupSession:
		res = GroupSession
	default:
		res = ""
	}

	if res == "" {
		return "", ErrInvalidSessionType
	}

	return res, nil
}
