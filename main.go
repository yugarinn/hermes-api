package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yugarinn/pigeon-api/http"
)


func setupRouter() *gin.Engine {
	router := gin.Default()
	routes.Register(router)

	return router
}
func main() {
	router := setupRouter()
	router.Run()
}
