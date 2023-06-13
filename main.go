package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yugarinn/hermes-api/http"
	"github.com/yugarinn/hermes-api/lib"
)


func setupRouter(hermes *lib.Hermes) *gin.Engine {
	router := gin.Default()
	routes.Register(hermes, router)

	return router
}

func main() {
	// TODO: finish the IoC Container implementation, think about maybe naming it `App` instead of `Hermes`.
    hermes := &lib.Hermes{
        dependencies: make(map[string]interface{}),
    }

	router := setupRouter(hermes)
	router.Run()
}
