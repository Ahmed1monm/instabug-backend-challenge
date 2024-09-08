package models

type Application struct {
	ID         uint64 `gorm:"primaryKey"`
	Token      string `gorm:"uniqueIndex:idx_token,length:255"`
	Name       string
	ChatsCount int64
	Chats      []Chat `gorm:"foreignKey:ApplicationID"`
}
