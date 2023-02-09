package breeds

import (
	"crypto/rand"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/yugarinn/pigeon-api/connections"
	"github.com/yugarinn/pigeon-api/app/messages/inputs"
	"github.com/yugarinn/pigeon-api/app/messages/models"
)


var database *gorm.DB = connections.Database()

func CreateMessage(data inputs.CreateMessageInput) (messages.Message, error) {
    message := messages.Message{Uuid: generateUuid(), Body: data.Body, SenderId: data.SenderId, ReceiverId: data.ReceiverId, DeliverAt: generateDeliveryDate()}
	result := database.Create(&message)

	return message, result.Error
}

func generateUuid() string {
	token := make([]byte, 32)
	rand.Read(token)

	return fmt.Sprintf("%x", token)
}

func generateDeliveryDate() time.Time {
	now := time.Now()

	return now.AddDate(0, 0, 1).UTC()
}
