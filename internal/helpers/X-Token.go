package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DecodeToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token = ctx.GetHeader("X-Token")
		if token == "" {
			ctx.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("u do not have the token"))
			return
		}
		var emailID, err = DecodeJWT(token)
		if err != nil {
			ctx.AbortWithError(http.StatusNotAcceptable, fmt.Errorf("token Error"))
			return
		}
		ctx.Set("Email", emailID)
		ctx.Next()
	}
}
