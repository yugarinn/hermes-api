package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yugarinn/hermes-api/app/users/factories"
	"github.com/yugarinn/hermes-api/app/users/inputs"
	"github.com/yugarinn/hermes-api/tests/mocks"
)


type ExpectedConfirmationResponse struct {
	ID uint64
	Email string
	Name string
	LastName string
	CountryCode string
	PhoneNumber string
	PhonePrefix string
}

func TestValidations(t *testing.T) {
	t.Run("POST /users/:userId/validations creates and sends a new user confirmation", func(t *testing.T) {
		ResetDatabase()

		factories.CreateUser(inputs.CreateUserInput{PhonePrefix: "34", PhoneNumber: "166666666", CountryCode: "ES"})

		app, router := SetupRouter()
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/users/1/validations", nil)

		router.ServeHTTP(writer, request)

		assert.Equal(t, 201, writer.Code)
		assert.Equal(t, true, DatabaseHas("users_validations", "user_id='1' AND is_used=0"))
		twilioMock := app.TwilioClient.(*mocks.TwilioMock)
		assert.Equal(t, 1, twilioMock.TimesInvoked)
	})

	t.Run("PUT /users/:userId/validations/:validationId validates the user and uses the validation if the code is correct", func(t *testing.T) {
		ResetDatabase()

		validationCode := "123ABC"

		user := factories.CreateUser(inputs.CreateUserInput{PhonePrefix: "34", PhoneNumber: "166666666", CountryCode: "ES"})
		factories.CreateUserValidation(factories.UserValidationFactoryInput{UserID: user.ID, Code: validationCode, IsUsed: false, ExpiresAt: time.Now().Add(24 * time.Hour)})

		var payload = []byte(`{"validationCode":"123ABC"}`)

		_, router := SetupRouter()
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/users/1/validations/1", bytes.NewBuffer(payload))

		router.ServeHTTP(writer, request)

		assert.Equal(t, 200, writer.Code)
		assert.Equal(t, true, DatabaseHas("users_validations", "id='1' AND user_id='1' AND is_used=1"))
		assert.Equal(t, true, DatabaseHas("users_users", "id='1' AND is_confirmed='1'"))
	})
}
