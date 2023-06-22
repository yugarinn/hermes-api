package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	core "github.com/yugarinn/hermes-api/core"
	inputs "github.com/yugarinn/hermes-api/app/users/inputs"
	services "github.com/yugarinn/hermes-api/app/users/services"
)


func CreateUserValidation(app *core.App, context *gin.Context) {
	userId, _ := strconv.ParseUint(context.Param("userId"), 10, 64)
	input := inputs.GetUserInput{UserID: userId}
	validate := validator.New()

	validationErrors := validate.Struct(&input)

	if validationErrors != nil {
		FailWithHttpCode(context, 422, validationErrors.Error())
		return
	}

	getUserResult := services.GetUser(input)
	services.SendValidationSMS(app, getUserResult.User)

	context.JSON(http.StatusCreated, nil)
}

func ValidateUser(context *gin.Context) {
	userId, _ := strconv.ParseUint(context.Param("userId"), 10, 64)
	validationId, _ := strconv.ParseUint(context.Param("validationId"), 10, 64)
	input := inputs.ValidateUserInput{UserID: userId, ValidationID: validationId}
	validate := validator.New()

	context.BindJSON(&input)
	validationErrors := validate.Struct(input)

	if validationErrors != nil {
		FailWithHttpCode(context, 422, validationErrors.Error())
		return
	}

	validateUserResult := services.ValidateUser(input)

	if validateUserResult.Success == false {
		FailWithHttpCode(context, 422, "the_provided_validation_code_is_invalid")
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
