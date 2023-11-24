package user

import (
	"errors"

	"github.com/oklog/ulid/v2"
)

var UserList = make(map[string]*User)
var ErrUserNotFound = errors.New("user not found")

func FindUser(id ulid.ULID) (*User, error) {
	user, ok := UserList[id.String()]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func AddToUserList(user *User) {
	UserList[user.Id.String()] = user

}

func RemoveUser(userId ulid.ULID) {
	delete(UserList, userId.String())
}
