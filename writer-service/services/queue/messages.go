package queue

import "time"

const (
	MessageCreate = "message:create"
	ChatCreate    = "chat:create"
)

type MessageCreatePayload struct {
	ID string `json:"id"`

	ApplicationID    string `json:"application_id"`
	ApplicationToken string `json:"application_token"`

	ChatID  string `json:"chat_id"`
	Number  int64  `json:"number"`
	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChatCreatePayload struct {
	ID string `json:"id"`

	ApplicationID string `json:"application_id"`
	Number        int64  `json:"number"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
