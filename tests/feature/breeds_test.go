package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yugarinn/catapi.cat/app/breeds/factories"
	"github.com/yugarinn/catapi.cat/app/breeds/inputs"
)

type ExpectedBreedResponse struct {
	ID uint64
	Email string
	Name string
	PhoneNumber string
}

func TestBreeds(t *testing.T) {
	t.Run("GET /breeds returns a paginated list of breeds", func(t *testing.T) {
		Reset()

		factories.CreateBreedsList(100, inputs.CreateBreedInput{})

		router := SetupRouter()

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/breeds", nil)
		router.ServeHTTP(writer, request)

		response := []string{}

		json.NewDecoder(writer.Body).Decode(&response)

		assert.Equal(t, 200, writer.Code)
		assert.Equal(t, 20, len(response))
	})
}
