package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yugarinn/catapi.cat/http/controllers"
	"github.com/yugarinn/catapi.cat/http/middlewares"
)

func Register(router *gin.Engine) {

	authorizedRoutes := router.Group("/")
	authorizedRoutes.Use(middlewares.CheckAccessToken)
	{
		authorizedRoutes.GET("/breeds", controllers.GetBreeds)
		authorizedRoutes.GET("/breeds/:breedId", controllers.GetBreed)
		authorizedRoutes.GET("/facts", controllers.GetFacts)
	}
}
