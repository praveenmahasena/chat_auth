package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/helpers"
	"github.com/praveenmahasena647/chat-app/internal/postgres"
)

func CreateAccount(gctx *gin.Context) {
	user := postgres.NewUserStruct()
	if err := gctx.Bind(user); err != nil {
		gctx.JSONP(http.StatusNotAcceptable, "Invalid input")
		return
	}

	if password, passwordErr := helpers.HashPassword(user.Password); passwordErr != nil {
		gctx.JSONP(http.StatusInternalServerError, "Error during password hashing")
		return
	} else {
		user.Password = password
	}
	if userInsertErr := user.InsertOne(gctx); userInsertErr != nil {
		gctx.JSONP(http.StatusNotAcceptable, "User Already Exists")
		return
	}

	JWT, JWTErr := helpers.GenerateJWT(user.Email)
	if JWTErr != nil {
		gctx.JSONP(http.StatusInternalServerError, "Error during generating JWT")
		return
	}
	gctx.JSONP(http.StatusCreated, JWT)
}
