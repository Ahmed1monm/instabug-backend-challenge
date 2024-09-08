package main

import (
	"log"
	"os"
	"worker/db"
	"worker/handlers"
	"worker/models"
	"worker/services"

	"github.com/hibiken/asynq"
)

const (
	MessageCreate = "message:create"
	ChatCreate    = "chat:create"
)

func main() {
	log.Println("Worker started")

	db.InitDB("instabug")
	db.Migrate(models.Application{}, models.Chat{}, models.Message{})

	services.SetupElasticSearch()

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: os.Getenv("REDIS_URL")},
		asynq.Config{
			Concurrency: 10,
		},
	)

	mux := asynq.NewServeMux()

	mux.HandleFunc(MessageCreate, handlers.HandleMessageCreate)
	mux.HandleFunc(ChatCreate, handlers.HandleChatCreate)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
