package queue

import (
	"os"

	"github.com/hibiken/asynq"
)

func SetupQueue() *asynq.Client {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS_URL")})
	return client
}

var Client = SetupQueue()
