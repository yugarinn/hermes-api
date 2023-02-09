package services

import (
	inputs "github.com/yugarinn/pigeon-api/app/users/inputs"
	managers "github.com/yugarinn/pigeon-api/app/users/managers"
	models "github.com/yugarinn/pigeon-api/app/users/models"
)


type CreateUserResult struct {
	User models.User
	Error error
}

func CreateUserAndSendValidationCode(input inputs.CreateUserInput) CreateUserResult {
	user, creationError := managers.CreateUser(input)

	if creationError == nil {
		go func() {
			createAndSendValidationCodeFor(user)
		}()
	}

	return CreateUserResult{User: user, Error: creationError}
}

func createAndSendValidationCodeFor(user models.User) (models.UserValidationCode, error) {
	validationCode, creationError := managers.CreateValidationCodeFor(user.ID)

	if creationError == nil {
		sendValidationSMSTo(user, validationCode)
	}

	return validationCode, creationError
}

// TODO
func sendValidationSMSTo(user models.User, validationCode models.UserValidationCode) error {
	return nil
}
