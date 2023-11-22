package user

import "github.com/oklog/ulid/v2"

type User struct {
	Id   ulid.ULID
	Name string
}

func MakeNewUser(name string) (User, error) {
	u := User{
		Id:   ulid.Make(),
		Name: name,
	}
	return u, nil
}
