package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primarykey"`
	Username string `gorm:"unique"`
}
type Chat struct {
	ID uint `gorm:"primaryKey"`

	Participants []User `gorm:"many2many:chat_participants;"`
	Messages     []Message
}
type Message struct {
	ID uint `gorm:"primaryKey"`

	ChatID uint
	Chat   Chat

	SenderID uint
	Sender   User

	Content string

	CreatedAt time.Time
}
