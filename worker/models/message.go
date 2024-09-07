package models

import "time"

type Message struct {
	ID     uint64 `gorm:"primaryKey"`
	ChatID uint64 `gorm:"index"`
	Number int64  `gorm:"uniqueIndex:message_number"`
	Body   string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
