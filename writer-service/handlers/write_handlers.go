package handlers

import (
	"net/http"
	"strconv"

	"writer-service/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Create chat
func CreateChat(c echo.Context) error {
	token := c.Param("token")

	var app models.Application
	if err := models.DB.First(&app, "token = ?", token).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Application not found"})
	}

	var chat models.Chat
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		var lastChat models.Chat
		if err := tx.Where("application_id = ?", app.ID).Order("number desc").First(&lastChat).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		chatNumber := lastChat.Number + 1

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

// Create message
func CreateMessage(c echo.Context) error {
	token := c.Param("token")
	chatNumber, err := strconv.ParseInt(c.Param("number"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid chat number"})
	}

	var app models.Application
	if err := models.DB.First(&app, "token = ?", token).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Application not found"})
	}

	var chat models.Chat
	if err := models.DB.First(&chat, "application_id = ? AND number = ?", app.ID, chatNumber).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Chat not found"})
	}

	var message models.Message
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		var lastMessage models.Message
		if err := tx.Where("chat_id = ?", chat.ID).Order("number desc").First(&lastMessage).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		messageNumber := lastMessage.Number + 1

		message = models.Message{
			ChatID: chat.ID,
			Number: messageNumber,
			Body:   c.FormValue("body"),
		}
		if err := tx.Create(&message).Error; err != nil {
			return err
		}

		if err := tx.Model(&chat).Update("messages_count", gorm.Expr("messages_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create message"})
	}

	return c.JSON(http.StatusCreated, map[string]int64{"number": message.Number})
}

