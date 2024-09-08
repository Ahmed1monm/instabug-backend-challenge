package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"worker/services"

	"github.com/hibiken/asynq"
)

type MessageCreatePayload struct {
	ID string `json:"id"`

	ApplicationID    string `json:"application_id"`
	ApplicationToken string `json:"application_token"`
	Number           int64  `json:"number"`

	ChatID  string `json:"chat_id"`
	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func HandleMessageCreate(ctx context.Context, t *asynq.Task) error {
	var p MessageCreatePayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	customElasticSearchKey := fmt.Sprintf("%s-%s-%s", p.ApplicationToken, p.ChatID, p.ID)
	services.Index(customElasticSearchKey, []byte(p.Content))

	return nil
}
