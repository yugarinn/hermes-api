package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExpectedMessageCreationResponse struct {
	ID uint64
	Uuid string
	SenderId int
	ReceiverId int
	Body string
	SentAt string
}

func TestMessages(t *testing.T) {
	t.Run("POST /messages successfuly creates a new message", func(t *testing.T) {
		t.Skip()
		ResetDatabase()

		var response ExpectedMessageCreationResponse

		_, router := SetupRouter()

		var payload = []byte(`{"senderId":"1", "receiverId":"2", "body":"long ass letter"`)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/messages", bytes.NewBuffer(payload))
		router.ServeHTTP(writer, request)

		json.NewDecoder(writer.Body).Decode(&response)

		assert.Equal(t, 201, writer.Code)
		assert.Equal(t, "1", response.SenderId)
		assert.Equal(t, "2", response.ReceiverId)
		assert.Equal(t, "long ass letter", response.Body)
		assert.Equal(t, "", response.SentAt)

		assert.Equal(t, true, DatabaseHas("messages_messages", "sender_id='1' AND receiver_id='2'"))
	})
}
