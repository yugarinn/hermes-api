package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yugarinn/pigeon-api/app/users/factories"
	"github.com/yugarinn/pigeon-api/app/users/inputs"
)

type ExpectedUserResponse struct {
	ID uint64
	Email string
	Name string
	PhoneNumber string
}

func TestUsers(t *testing.T) {
	t.Run("GET /users/:userId returns the requested user", func(t *testing.T) {
		Reset()

		var response ExpectedUserResponse

		user := factories.CreateUser(inputs.CreateUserInput{Email: "foo@bar.com"})

		router := SetupRouter()

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/users/1", nil)
		AuthenticateAs(1 /* just pass the raw ID and save a DB call */, request)
		router.ServeHTTP(writer, request)

		json.NewDecoder(writer.Body).Decode(&response)

		assert.Equal(t, 200, writer.Code)
		assert.Equal(t, user.ID, response.ID)
		assert.Equal(t, user.Email, response.Email)
		assert.Equal(t, user.Name, response.Name)
		assert.Equal(t, user.PhoneNumber, response.PhoneNumber)
	})

	t.Run("POST /users creates a new user", func(t *testing.T) {
		Reset()

		var response ExpectedUserResponse

		router := SetupRouter()

		var payload = []byte(`{"email":"foo@bar.com", "name":"A Nice Name", "phoneNumber":"123456789", "countryCode":"ES"}`)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
		router.ServeHTTP(writer, request)

		json.NewDecoder(writer.Body).Decode(&response)

		assert.Equal(t, 201, writer.Code)
		assert.Equal(t, "foo@bar.com", response.Email)
		assert.Equal(t, "A Nice Name", response.Name)
		assert.Equal(t, "123456789", response.PhoneNumber)

		assert.Equal(t, true, DatabaseHas("users_users", "email='foo@bar.com' AND name='A Nice Name' AND phone_number='123456789' AND country_code='ES'"))
	})

	t.Run("POST /users creates a new password alongside the user", func(t *testing.T) {
		Reset()

		var response ExpectedUserResponse

		router := SetupRouter()

		var payload = []byte(`{"email":"foo@bar.com", "name":"A Nice Name", "phoneNumber":"123456789", "countryCode":"ES"}`)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
		router.ServeHTTP(writer, request)

		json.NewDecoder(writer.Body).Decode(&response)

		// FIXME: manually giving the goroutine time to create the validation
		time.Sleep(1 * time.Second)

		assert.Equal(t, 201, writer.Code)
		assert.Equal(t, true, DatabaseHas("users_passwords", "user_id=1"))
	})

	t.Run("POST /users does not create a new user without required fields", func(t *testing.T) {
		Reset()

		var response ExpectedUserResponse

		router := SetupRouter()

		var payload = []byte(`{"name":"A Nice Name", "phoneNumber":"123456789"}`)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
		router.ServeHTTP(writer, request)

		json.NewDecoder(writer.Body).Decode(&response)

		assert.Equal(t, 422, writer.Code)
		assert.Equal(t, true, DatabaseMissing("users_users", "name='A Nice Name' AND phone_number='123456789'"))
	})

	t.Run("POST /users does not create a new user if the provided phone_number already exists", func(t *testing.T) {
		Reset()

		repeatedEmail := "john@doe.com"

		factories.CreateUser(inputs.CreateUserInput{Email: repeatedEmail})

		var response ExpectedUserResponse

		router := SetupRouter()

		var payload = []byte(`{"email":"john@doe.com", "name":"John Doe", "phoneNumber":"123456789"}`)

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
		router.ServeHTTP(writer, request)

		json.NewDecoder(writer.Body).Decode(&response)

		assert.Equal(t, 422, writer.Code)
		assert.Equal(t, true, DatabaseHasCount("users_users", 1, "email='john@doe.com'"))
	})

	t.Run("GET /users/:userId does not allow to access another user's profile", func(t *testing.T) {
		Reset()

		var response ExpectedUserResponse

		factories.CreateUsersList(2, inputs.CreateUserInput{Email: "foo@bar.com"})

		router := SetupRouter()

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/users/2", nil)
		AuthenticateAs(1 /* just pass the raw ID and save a DB call */, request)
		router.ServeHTTP(writer, request)

		json.NewDecoder(writer.Body).Decode(&response)

		assert.Equal(t, 401, writer.Code)
	})
}
