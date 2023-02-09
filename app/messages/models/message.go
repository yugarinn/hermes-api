package messages

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Uuid string
	SenderId int
	ReceiverId int
	Body string
	DeliverAt time.Time
	DeliveredAt time.Time
}

func (Message) TableName() string {
	return "messages_message"
}
