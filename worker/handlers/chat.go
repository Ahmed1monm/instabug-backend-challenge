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

	"github.com/hibiken/asynq"
)

type ChatCreatePayload struct {
	ID string `json:"id"`

	ApplicationID string `json:"application_id"`
	ChatID        string `json:"chat_id"`
	Number        int64  `json:"number"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func HandleChatCreate(ctx context.Context, t *asynq.Task) error {
	var p ChatCreatePayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	applicationId, err := strconv.ParseUint(p.ApplicationID, 10, 64)

	if err != nil {
		return fmt.Errorf("strconv.ParseInt failed: %v: %w", err, asynq.SkipRetry)
	}

	id, err := strconv.ParseUint(p.ID, 10, 64)

	if err != nil {
		return fmt.Errorf("strconv.ParseInt failed: %v: %w", err, asynq.SkipRetry)
	}

	tx := db.DB.Create(&models.Chat{
		ID:            id,
		ApplicationID: applicationId,
		Number:        p.Number,
	})

	if tx.Error != nil {
		log.Printf("Chat creation failed: %v", tx.Error)
		return fmt.Errorf("db.DB.Create failed: %v: %w", tx.Error, asynq.SkipRetry)
	}

	return nil
}
