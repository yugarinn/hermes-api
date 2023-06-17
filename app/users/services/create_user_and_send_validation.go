package services

import (
	"fmt"

	core "github.com/yugarinn/hermes-api/core"
	inputs "github.com/yugarinn/hermes-api/app/users/inputs"
	managers "github.com/yugarinn/hermes-api/app/users/managers"
	models "github.com/yugarinn/hermes-api/app/users/models"
	utils "github.com/yugarinn/hermes-api/utils"
)


type CreateUserResult struct {
	User models.User
	Error error
}

func CreateUserAndSendValidationCode(app *core.App, input inputs.CreateUserInput) CreateUserResult {
	user, creationError := managers.CreateUser(input)

	if creationError == nil {
		if utils.IsProduction() {
			go sendSMS(app, user)
		} else {
			sendSMS(app, user)
		}
	}

	return CreateUserResult{User: user, Error: creationError}
}

func sendSMS(app *core.App, user models.User) {
	validationCode, creationError := managers.CreateValidationCodeFor(user.ID)

	if creationError == nil {
		toPhoneNumber := fmt.Sprintf("%s%s", user.PhonePrefix, user.PhoneNumber)
		fromPhoneNumber := "+34667888999"

		app.TwilioClient.SendSMS(toPhoneNumber, fromPhoneNumber, validationCode.Code)
	}
}
