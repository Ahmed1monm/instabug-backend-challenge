package models

import "time"

type Chat struct {
	ID            uint  `gorm:"primaryKey"`
	ApplicationID uint  `gorm:"index"`
	Number        int64 `gorm:"uniqueIndex:chat_number"`
	MessagesCount int64
	Messages      []Message `gorm:"foreignKey:ChatID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
