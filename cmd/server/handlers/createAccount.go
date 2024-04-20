package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/cmd/helpers"
	"github.com/praveenmahasena647/chat-app/cmd/postgres"
)

func CreateAccount(gctx *gin.Context) {
	var user = postgres.NewUser()
	if err := gctx.ShouldBind(&user); err != nil {
		gctx.JSONP(http.StatusNotAcceptable, "Input not allowed")
		return
	}
	if h, e := helpers.HashPassword(user.Password); e != nil {
		gctx.JSONP(http.StatusInternalServerError, "Password Hash Error")
		return
	} else {
		user.Password = h
	}
	var ctx, cancel = context.WithTimeout(context.Background(), time.Minute*20)
	defer cancel()

	if err := user.Insert(ctx); err != nil {

	} else {

	}

}
