package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/helpers"
	"github.com/praveenmahasena647/chat-app/internal/postgres"
)

func MarkVerify(ctx *gin.Context) {
	id := ctx.Param("id")
	email, decodeErr := helpers.DecodeJWTVerify(id)
	if decodeErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, decodeErr.Error())
		return
	}
	verified, verifyErr := postgres.Verify(ctx, email)
	if verifyErr != nil {
		ctx.JSONP(http.StatusNotFound, verifyErr)
		return
	}
	if !verified {
		ctx.JSONP(http.StatusInternalServerError, errors.New("error during changing the Status"))
		return
	}
	ctx.JSONP(http.StatusOK, "Verified")
}
