package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"worker/db"
	"worker/models"
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

	id, err := strconv.ParseUint(p.ID, 10, 64)

	if err != nil {
		return fmt.Errorf("strconv.ParseInt failed: %v: %w", err, asynq.SkipRetry)
	}

	chatId, err := strconv.ParseUint(p.ChatID, 10, 64)

	if err != nil {
		return fmt.Errorf("strconv.ParseInt failed: %v: %w", err, asynq.SkipRetry)
	}

	tx := db.DB.Create(&models.Message{
		ID:     uint64(id),
		ChatID: uint64(chatId),
		Body:   p.Content,
		Number: p.Number,
	})

	customElasticSearchKey := fmt.Sprintf("%s-%s-%s", p.ApplicationToken, p.ChatID, p.ID)
	services.Index(customElasticSearchKey, []byte(p.Content))

	if tx.Error != nil {
		log.Printf("Message creation failed: %v", tx.Error)
		return fmt.Errorf("db.DB.Create failed: %v: %w", tx.Error, asynq.SkipRetry)
	}

	return nil
}
