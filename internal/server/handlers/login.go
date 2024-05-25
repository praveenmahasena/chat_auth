package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/helpers"
	"github.com/praveenmahasena647/chat-app/internal/postgres"
)

func Login(ctx *gin.Context) {
	u := postgres.NewLoginCredentials()
	if err := ctx.Bind(u); err != nil {
		ctx.JSONP(http.StatusMethodNotAllowed, "the credentials are not allowed")
		return
	}
	usrInfo, usrInfoErr := u.Login(ctx)
	if usrInfoErr != nil {
		ctx.JSONP(http.StatusNotFound, usrInfoErr)
		return
	}
	if err := helpers.VerifyPassword(usrInfo.Password, u.Password); err != nil {
		ctx.JSONP(http.StatusNotFound, err)
		return
	}
	JWT, JWTErr := helpers.GenerateJWT(u.EmailID)
	if JWTErr != nil {
		ctx.JSONP(http.StatusNotFound, JWTErr)
		return
	}
	ctx.JSONP(http.StatusOK, JWT)
}
