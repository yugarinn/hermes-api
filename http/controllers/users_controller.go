package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	inputs "github.com/yugarinn/pigeon-api/app/users/inputs"
	services "github.com/yugarinn/pigeon-api/app/users/services"
	responses "github.com/yugarinn/pigeon-api/http/responses"
)


func CreateUser(context *gin.Context) {
	input := inputs.CreateUserInput{}
	validate := validator.New()

	context.BindJSON(&input)
	validationErrors := validate.Struct(input)

	if  validationErrors != nil {
		HttpErrorHandler(context, 422, validationErrors.Error())
		return
	}

	result := services.CreateUserAndSendValidationCode(input)

	if result.Error != nil {
		HttpErrorHandler(context, 422, result.Error.Error())
		return
	}

	context.JSON(http.StatusCreated, responses.SerializeUser(result.User))
}
