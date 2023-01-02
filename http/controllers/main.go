package controllers

import (
	"github.com/gin-gonic/gin"
)


const DEFAULT_PAGE_SIZE = 20

func HttpErrorHandler(context *gin.Context, code int, message string) {
	context.AbortWithStatusJSON(code, gin.H{"message": message})
}
