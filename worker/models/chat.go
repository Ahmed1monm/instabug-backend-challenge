package models

import "time"

type Chat struct {
	ID            uint64 `gorm:"primaryKey"`
	ApplicationID uint64 `gorm:"index"`
	Number        int64  `gorm:"uniqueIndex:chat_number;autoIncrement"`

	MessagesCount int64
	Messages      []Message `gorm:"foreignKey:ChatID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
