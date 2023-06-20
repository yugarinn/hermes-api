package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	core "github.com/yugarinn/hermes-api/core"
	inputs "github.com/yugarinn/hermes-api/app/users/inputs"
	services "github.com/yugarinn/hermes-api/app/users/services"
	responses "github.com/yugarinn/hermes-api/http/responses"
)


func CreateUser(app *core.App, context *gin.Context) {
	input := inputs.CreateUserInput{}
	validate := validator.New()

	context.BindJSON(&input)
	validationErrors := validate.Struct(input)

	if validationErrors != nil {
		FailWithHttpCode(context, 422, validationErrors.Error())
		return
	}

	result := services.CreateUserAndSendValidationCode(app, input)

	if result.Error != nil {
		FailWithHttpCode(context, 422, result.Error.Error())
		return
	}

	context.JSON(http.StatusCreated, responses.SerializeUser(result.User))
}

func GetPaginatedUsers(context *gin.Context) {
	input := inputs.GetPaginatedUsersInput{Page: 1, Size: 10}
	result := services.GetPaginatedUsers(input)

	users := []responses.UserResponse{}

	for _, user := range result.Users {
		users = append(users, responses.SerializeUser(user))
	}

	response := responses.PaginatedUserResponse{Users: users, Pagination: responses.Pagination{Page: 1, Size: 10, Total: 10 }}

	context.JSON(http.StatusOK, response)
}

func GetUser(context *gin.Context) {
	userId, _ := strconv.ParseUint(context.Param("userId"), 10, 64)
	input := inputs.GetUserInput{UserID: userId}

	// if GetAuthenticatedUser(context).ID != userId {
	//     context.AbortWithError(401, errors.New("unauthorized"))
	// }

	result := services.GetUser(input)

	if result.Error != nil {
		FailWithHttpCode(context, 404, "user not found")
		return
	}

	context.JSON(http.StatusOK, responses.SerializeUser(result.User))
}
