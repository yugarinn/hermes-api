package services

import (
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"

	inputs "github.com/yugarinn/hermes-api/app/users/inputs"
	managers "github.com/yugarinn/hermes-api/app/users/managers"
	models "github.com/yugarinn/hermes-api/app/users/models"
)


type CreateUserResult struct {
	User models.User
	Error error
}

func CreateUserAndSendValidationCode(input inputs.CreateUserInput) CreateUserResult {
	user, creationError := managers.CreateUser(input)

	if creationError == nil {
		go func() {
			CreateAndSendValidationCodeFor(user)
		}()
	}

	return CreateUserResult{User: user, Error: creationError}
}

func CreateAndSendValidationCodeFor(user models.User) (models.UserValidationCode, error) {
	validationCode, creationError := managers.CreateValidationCodeFor(user.ID)

	if creationError == nil {
		sendValidationSMSTo(user, validationCode)
	}

	return validationCode, creationError
}

func sendValidationSMSTo(user models.User, validationCode models.UserValidationCode) error {
	client := twilio.NewRestClient()

	params := &twilioApi.CreateMessageParams{}

	params.SetBody(fmt.Sprintf("This is your verification code: %s", validationCode.Code))
	params.SetFrom("+12765229854")
	params.SetTo(fmt.Sprintf("%s%s", user.PhonePrefix, user.PhoneNumber))

	_, error := client.Api.CreateMessage(params)

	return error
}
