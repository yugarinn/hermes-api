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
	FromId int
	ToId int
	Body string
	SentAt string
}

func TestBreeds(t *testing.T) {
	t.Run("POST /messages successfuly creates a new message", func(t *testing.T) {
		Reset()

		var response ExpectedMessageCreationResponse

		router := SetupRouter()

		var payload = []byte(`{"fromId":"1", "toId":"2", "body":"long ass letter"`)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
		router.ServeHTTP(writer, request)

		json.NewDecoder(writer.Body).Decode(&response)

		assert.Equal(t, 201, writer.Code)
		assert.Equal(t, "1", response.FromId)
		assert.Equal(t, "2", response.ToId)
		assert.Equal(t, "long ass letter", response.Body)
		assert.Equal(t, "", response.SentAt)

		assert.Equal(t, true, DatabaseHas("messages_messages", "from_id='1' AND to_id='2'"))
	})
}
