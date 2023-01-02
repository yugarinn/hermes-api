package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// TODO
func CheckAccessToken(context *gin.Context) {
    providedToken := context.GetHeader("Access-Token")

    if false && providedToken == "" {
        context.AbortWithError(401, errors.New("unauthorized"))
        return
    } else {
		context.Next()
	}

}
