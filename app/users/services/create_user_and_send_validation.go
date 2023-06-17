package services

import (
	"fmt"

	core "github.com/yugarinn/hermes-api/core"
	inputs "github.com/yugarinn/hermes-api/app/users/inputs"
	managers "github.com/yugarinn/hermes-api/app/users/managers"
	models "github.com/yugarinn/hermes-api/app/users/models"
)


type CreateUserResult struct {
	User models.User
	Error error
}

func CreateUserAndSendValidationCode(app *core.App, input inputs.CreateUserInput) CreateUserResult {
	user, creationError := managers.CreateUser(input)

	if creationError == nil {
		go func() {
			validationCode, creationError := managers.CreateValidationCodeFor(user.ID)

			if creationError == nil {
				toPhoneNumber := fmt.Sprintf("%s%s", user.PhonePrefix, user.PhoneNumber)
				fromPhoneNumber := "+34667888999"

				app.TwilioClient.SendSMS(toPhoneNumber, fromPhoneNumber, validationCode.Code)
			}

		}()
	}

	return CreateUserResult{User: user, Error: creationError}
}
