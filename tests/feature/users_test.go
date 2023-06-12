package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yugarinn/hermes-api/app/users/factories"
	"github.com/yugarinn/hermes-api/app/users/inputs"
)


type ExpectedUserResponse struct {
	ID uint64
	Email string
	Name string
	LastName string
	CountryCode string
	PhoneNumber string
	PhonePrefix string
}

func TestUsers(t *testing.T) {
	t.Run("POST /users creates a new user", func(t *testing.T) {
		ResetDatabase()

		var response ExpectedUserResponse

		router := SetupRouter()

		var payload = []byte(`{"phonePrefix":"+34", "phoneNumber":"123456789", "countryCode":"ES"}`)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
		router.ServeHTTP(writer, request)

		json.NewDecoder(writer.Body).Decode(&response)

		assert.Equal(t, 201, writer.Code)
		assert.Equal(t, "+34", response.PhonePrefix)
		assert.Equal(t, "123456789", response.PhoneNumber)
		assert.Equal(t, "ES", response.CountryCode)

		assert.Equal(t, true, DatabaseHas("users_users", "phone_prefix='+34' AND phone_number='123456789' AND country_code='ES'"))
	})

	t.Run("GET /users/:userId returns the requested user", func(t *testing.T) {
		ResetDatabase()

		var response ExpectedUserResponse

		user := factories.CreateUser(inputs.CreateUserInput{PhonePrefix: "34", PhoneNumber: "666666666", CountryCode: "ES"})

		router := SetupRouter()

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/users/1", nil)
		AuthenticateAs(1 /* just pass the raw ID and skip the DB call */, request)
		router.ServeHTTP(writer, request)

		json.NewDecoder(writer.Body).Decode(&response)

		fmt.Println(response.PhoneNumber)
		fmt.Println(user.PhoneNumber)

		assert.Equal(t, 200, writer.Code)
		assert.Equal(t, user.ID, response.ID)
		assert.Equal(t, user.Name, response.Name)
		assert.Equal(t, user.LastName, response.LastName)
		assert.Equal(t, user.PhonePrefix, response.PhonePrefix)
		assert.Equal(t, user.PhoneNumber, response.PhoneNumber)
		assert.Equal(t, user.CountryCode, response.CountryCode)
	})
}
