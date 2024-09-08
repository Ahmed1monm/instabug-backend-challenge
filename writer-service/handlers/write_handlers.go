package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"writer-service/db"
	"writer-service/models"
	"writer-service/services/queue"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func CreateChat(c echo.Context) error {
	token := c.Param("token")

	var app models.Application

	if err := db.DB.First(&app, "token = ?", token).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Application not found"})
	}

	var chat models.Chat
	var chatNumber int64

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		var lastChat models.Chat

		if err := tx.Where("application_id = ?", app.ID).
			Order("number desc").
			Limit(1).
			First(&lastChat).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		chatNumber = lastChat.Number + 1

		chat = models.Chat{
			ApplicationID: app.ID,
			Number:        chatNumber,
		}
		if err := tx.Create(&chat).Error; err != nil {
			return err
		}

		if err := tx.Model(&app).Update("chats_count", gorm.Expr("chats_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create chat"})
	}

	return c.JSON(http.StatusCreated, map[string]int64{"number": chat.Number})
}

func CreateMessage(c echo.Context) error {
	token := c.Param("token")
	chatNumber, err := strconv.ParseInt(c.Param("number"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid chat number",
		})
	}

	type RequestPayload struct {
		Body string `json:"body"`
	}

	var payload RequestPayload

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	var app models.Application

	if err := db.DB.First(&app, "token = ?", token).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Application not found"})
	}

	var chat models.Chat
	if err := db.DB.First(&chat, "application_id = ? AND number = ?", app.ID, chatNumber).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Chat not found"})
	}

	var message models.Message
	var messageNumber int64

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		var lastMessage models.Message
		if err := tx.Where("chat_id = ?", chat.ID).Order("number desc").First(&lastMessage).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		messageNumber = lastMessage.Number + 1

		message = models.Message{
			ChatID: uint(chat.ID),
			Number: messageNumber,
			Body:   payload.Body,
		}
		if err := tx.Create(&message).Error; err != nil {
			return err
		}

		if err := tx.Model(&chat).Update("messages_count", gorm.Expr("messages_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})

	customId := fmt.Sprintf("%s_%d_%d", app.Token, message.ChatID, message.ID)

	task, err := queue.NewMessageCreateTask(queue.MessageCreatePayload{
		ID:        customId,
		ChatID:    strconv.FormatUint(uint64(chat.ID), 10),
		Content:   message.Body,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create message"})
	}

	info, err := queue.Client.Enqueue(task, asynq.MaxRetry(3), asynq.Timeout(1*time.Minute))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create message"})
	}

	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"content": message.Body,
		"number":  messageNumber,
	})
}
