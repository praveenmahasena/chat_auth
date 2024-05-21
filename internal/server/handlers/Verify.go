package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/helpers"
	"github.com/praveenmahasena647/chat-app/internal/postgres"
)

func Verify(gctx *gin.Context) {
	email, _ := gctx.Get("Email")
	verified, verifiedErr := postgres.IsVerified(gctx, email.(string))
	if verifiedErr != nil {
		gctx.JSONP(http.StatusNotFound, "user does not exists")
		return
	}
	if verified {
		gctx.JSONP(http.StatusMethodNotAllowed, "user Already verified")
		return
	}
	JWT, JWTErr := helpers.GenerateJWTVerify(email.(string))
	if JWTErr != nil {
		gctx.JSONP(http.StatusInternalServerError, "couldnt make JWT")
		return
	}
	mailErr := sendMail(JWT, email.(string))
	if mailErr != nil {
		return
	}
	gctx.JSONP(http.StatusOK, "mail sent")
}

func sendMail(JWT, mailID string) error {
	return nil
}
