package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yugarinn/hermes-api/http/controllers"
	"github.com/yugarinn/hermes-api/http/middlewares"
	"github.com/yugarinn/hermes-api/lib"
)

func Register(hermes *lib.Hermes, router *gin.Engine) {

	router.POST("/users", controllers.CreateUser)

	authorizedRoutes := router.Group("/")
	authorizedRoutes.Use(middlewares.CheckAccessToken)
	{
		authorizedRoutes.POST("/messages", controllers.CreateMessage)
		authorizedRoutes.GET("/users/:userId", controllers.GetUser)
	}
}
