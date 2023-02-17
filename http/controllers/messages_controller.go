package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	inputs "github.com/yugarinn/pigeon-api/app/messages/inputs"
	managers "github.com/yugarinn/pigeon-api/app/messages/managers"
	responses "github.com/yugarinn/pigeon-api/http/responses"
)


func CreateMessage(context *gin.Context) {
	input := inputs.CreateMessageInput{}

	validate := validator.New()

	context.BindJSON(&input)
	validationErrors := validate.Struct(input)

	if  validationErrors != nil {
		FailWithHttpCode(context, 422, validationErrors.Error())
		return
	}

	message, creationError := managers.CreateMessage(input)

	if creationError != nil {
		FailWithHttpCode(context, 422, creationError.Error())
		return
	}

	context.JSON(http.StatusCreated, responses.SerializeMessage(message))
}
