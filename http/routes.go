package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yugarinn/hermes-api/http/controllers"
	"github.com/yugarinn/hermes-api/http/middlewares"
	"github.com/yugarinn/hermes-api/core"
)

func Register(app *core.App, router *gin.Engine) {
    router.POST("/users", func(context *gin.Context) { controllers.CreateUser(app, context) })

    router.POST("/users/:userId/validations", func(context *gin.Context) { controllers.CreateUserValidation(app, context) })
    router.PUT("/users/:userId/validations/:validationId", controllers.ValidateUser)

	authorizedRoutes := router.Group("/")
	authorizedRoutes.Use(middlewares.CheckAccessToken)
	{
		authorizedRoutes.POST("/messages", controllers.CreateMessage)
		authorizedRoutes.GET("/users/:userId", controllers.GetUser)
	}
}
