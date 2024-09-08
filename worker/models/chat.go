package models

import "time"

type Chat struct {
	ID uint64 `gorm:"primaryKey"`

	MessagesCount int64
	Messages      []Message `gorm:"foreignKey:ChatID"`
	Number        int64     `gorm:"uniqueIndex:idx_app_chat_number,priority:2"`
	ApplicationID uint64    `gorm:"uniqueIndex:idx_app_chat_number,priority:1"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
