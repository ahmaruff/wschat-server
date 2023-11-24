package wsservice

import (
	"ahmaruff/wschat/user"
	"errors"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

const PrivateSession = "PRIVATE"
const GroupSession = "GROUP"

var ErrInvalidSessionType = errors.New("chat: invalid session type")

type Session struct {
	Id           ulid.ULID
	Type         string
	LastActive   time.Time
	Participants map[ulid.ULID]*user.User //store id user
}

func MakeNewSession(id ulid.ULID, t string) (Session, error) {
	new_t, err := validateType(t)

	if err != nil {
		return Session{}, err
	}

	s := Session{
		Id:           id,
		Type:         new_t,
		LastActive:   time.Now(),
		Participants: make(map[ulid.ULID]*user.User),
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
