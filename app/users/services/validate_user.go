package services

import (
	"time"

	inputs "github.com/yugarinn/hermes-api/app/users/inputs"
	managers "github.com/yugarinn/hermes-api/app/users/managers"
)


type ValidateUserResult struct {
	Success bool
	Error error
}

func ValidateUser(input inputs.ValidateUserInput) ValidateUserResult {
	userValidation, userValidationRetrievalError := managers.GetUserValidation(input.ValidationID)

	if userValidationRetrievalError != nil {
		return ValidateUserResult{Success: false, Error: userValidationRetrievalError}
	}

	validationBelongsToUser := userValidation.UserId == input.UserID
	validationIsNotUsed := userValidation.IsUsed == false
	validationIsNotExpired := userValidation.ExpiresAt.After(time.Now())
	providedCodeEqualsTheValidationCode := userValidation.Code == input.ValidationCode

	if validationBelongsToUser && validationIsNotUsed && validationIsNotExpired && providedCodeEqualsTheValidationCode {
		managers.SetUserAsConfirmed(input.UserID)
		managers.SetValidationAsUsed(input.ValidationID)

		return ValidateUserResult{Success: true, Error: nil}
	} else {
		return ValidateUserResult{Success: false, Error: nil}
	}
}
