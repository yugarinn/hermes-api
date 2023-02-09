package inputs

type CreateMessageInput struct {
	Uuid string `json:"uuid"`
	SenderId int `json:"senderId" validate:"required"`
	ReceiverId int `json:"receiverId" validate:"required"`
	Body string `json:"body" validate:"required"`
}
