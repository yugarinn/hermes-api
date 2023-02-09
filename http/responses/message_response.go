package responses

import (
	models "github.com/yugarinn/pigeon-api/app/messages/models"
)

type MessageResponse struct {
	ID string          `json:"id"`
	Body string        `json:"body"`
}

func SerializeMessage(message models.Message) MessageResponse {
	return MessageResponse{
		ID: message.Uuid,
		Body: message.Body,
	}
}
