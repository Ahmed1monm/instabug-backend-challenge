package models

type Application struct {
	ID         uint   `gorm:"primaryKey"`
	Token      string `gorm:"uniqueIndex"`
	Name       string
	ChatsCount int64
	Chats      []Chat `gorm:"foreignKey:ApplicationID"`
}

type Chat struct {
	ID            uint   `gorm:"primaryKey"`
	ApplicationID uint   `gorm:"index"`
	Number        int64  `gorm:"uniqueIndex:chat_number"`
	MessagesCount int64
	Messages      []Message `gorm:"foreignKey:ChatID"`
}

type Message struct {
	ID     uint   `gorm:"primaryKey"`
	ChatID uint   `gorm:"index"`
	Number int64  `gorm:"uniqueIndex:message_number"`
	Body   string `gorm:"type:text"`
}