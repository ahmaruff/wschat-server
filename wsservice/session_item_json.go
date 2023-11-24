package wsservice

import (
	"ahmaruff/wschat/user"
	"encoding/json"
	"time"

	"github.com/oklog/ulid/v2"
)

func (s Session) MarshalJSON() ([]byte, error) {
	var j struct {
		Id           ulid.ULID                `json:"id"`
		Type         string                   `json:"type"`
		LastActive   time.Time                `json:"last_active"`
		Participants map[ulid.ULID]*user.User `json:"participants"`
	}

	j.Id = s.Id
	j.Type = s.Type
	j.LastActive = s.LastActive
	j.Participants = s.Participants

	return json.Marshal(j)
}

func (s *Session) UnmarshalJSON(data []byte) error {
	var j struct {
		Id           ulid.ULID                `json:"id"`
		Type         string                   `json:"type"`
		LastActive   time.Time                `json:"last_active"`
		Participants map[ulid.ULID]*user.User `json:"participants"`
	}

	err := json.Unmarshal(data, &j)

	if err != nil {
		return err
	}

	s = &Session{
		Id:           j.Id,
		Type:         j.Type,
		LastActive:   j.LastActive,
		Participants: j.Participants,
	}

	return nil
}
