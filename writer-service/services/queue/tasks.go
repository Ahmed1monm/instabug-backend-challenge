package queue

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

func NewMessageCreateTask(payload MessageCreatePayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(MessageCreate, data), nil
}

func NewChatCreateTask(payload ChatCreatePayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(ChatCreate, data), nil
}
