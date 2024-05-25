package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/helpers"
)

func RouteGurd(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")
	if token == "" {
		ctx.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("u do not have the token"))
		return
	}
	emailID, err := helpers.DecodeJWT(token)
	if err != nil {
		ctx.AbortWithError(http.StatusNotAcceptable, fmt.Errorf("token Error"))
		return
	}
	ctx.Set("Email", emailID)
	ctx.Next()
}
