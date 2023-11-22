package chat

import "github.com/oklog/ulid/v2"

type Participant struct {
	ChatId ulid.ULID
	UserId ulid.ULID
}
