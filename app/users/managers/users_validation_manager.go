package users

import (
	"math/rand"
	"strconv"
	"time"

	users "github.com/yugarinn/hermes-api/app/users/models"
)


const validationCodeLifetimeInSeconds = 120

func CreateValidationCodeFor(userId uint64) (users.UserValidation, error) {
    validation := users.UserValidation{UserId: userId, Code: generateValidationCode(), IsUsed: false, ExpiresAt: generateExpirationDate()}
	result := database.Create(&validation)

	return validation, result.Error
}

func GetUserValidation(validationId uint64) (users.UserValidation, error) {
	var validation users.UserValidation
	result := database.Where("id", validationId).First(&validation)

	return validation, result.Error
}

func SetValidationAsUsed(id uint64) (users.UserValidation, error) {
	validation, retrievalError := GetUserValidation(id)

	if retrievalError != nil {
		return users.UserValidation{}, retrievalError
	}

	validation.IsUsed = true
	updateResult := database.Save(&validation)

	return validation, updateResult.Error
}

func generateValidationCode() string {
	code := strconv.Itoa(rand.Intn(1000000))

	// TODO: make sure there are no active `UserValidationCode` entries with the same `code`

	return code
}

// TODO: the expiration date should be 2 minutes from now.
// Use `validationCodeLifetimeInSeconds`.
func generateExpirationDate() time.Time {
	now := time.Now()

	return now.AddDate(0, 0, 1).UTC()
}
