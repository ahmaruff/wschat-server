package user

import (
	"encoding/json"

	"github.com/oklog/ulid/v2"
)

func (t User) MarshalJSON() ([]byte, error) {
	var j struct {
		Id   ulid.ULID `json:"id"`
		Name string    `json:"name"`
	}
	j.Id = t.Id
	j.Name = t.Name

	return json.Marshal(j)
}

func (u *User) UnmarshalJSON(data []byte) error {
	var j struct {
		Id   ulid.ULID `json:"id"`
		Name string    `json:"name"`
	}

	err := json.Unmarshal(data, &j)

	if err != nil {
		return err
	}

	u = &User{
		Id:   j.Id,
		Name: j.Name,
	}

	return nil
}
