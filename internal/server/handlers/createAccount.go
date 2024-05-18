package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/helpers"
	"github.com/praveenmahasena647/chat-app/internal/postgres"
)

func CreateAccount(gctx *gin.Context) {
	var usr = new(postgres.User)
	if err := gctx.Bind(usr); err != nil {
		gctx.JSONP(http.StatusNotAcceptable, "Invalid input")
		return
	}

	if p, err := helpers.HashPassword(usr.Password); err != nil {
		gctx.JSONP(http.StatusInternalServerError, "Error during password hashing")
		return
	} else {
		usr.Password = p
	}
	if err := usr.InsertOne(gctx); err != nil {
		gctx.JSONP(http.StatusNotAcceptable, "User Already Exists")
		return
	}

	var jwt, jwtErr = helpers.GenerateJWT(usr.Email)
	if jwtErr != nil {
		gctx.JSONP(http.StatusInternalServerError, "Error during generating JWT")
		return
	}
	gctx.JSONP(http.StatusCreated, jwt)
}
