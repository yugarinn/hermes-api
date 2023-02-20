package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yugarinn/pigeon-api/app/users/factories"
	"github.com/yugarinn/pigeon-api/app/users/inputs"
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
	t.Run("GET /users/:userId returns the requested user", func(t *testing.T) {
		Reset()

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
