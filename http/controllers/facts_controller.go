package controllers

import (
	"github.com/gin-gonic/gin"
)


func GetFacts(context *gin.Context) {
	context.AbortWithStatusJSON(422, "")
}
