package chat

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Message struct {
	Id        ulid.ULID
	SessionId ulid.ULID
	SenderId  ulid.ULID
	Message   string
	CreatedAt time.Time
}
type MessageData struct {
	SessionId string
	SenderId  string
	Message   string
}

func MakeNewMessage(sessionId, senderId, m string) (Message, error) {
	session_id, err := ulid.Parse(sessionId)
	if err != nil {
		return Message{}, err
	}

	sender_id, err := ulid.Parse(senderId)
	if err != nil {
		return Message{}, err
	}

	mes := Message{
		Id:        ulid.Make(),
		SessionId: session_id,
		SenderId:  sender_id,
		Message:   m,
		CreatedAt: time.Now(),
	}

	return mes, nil
}
