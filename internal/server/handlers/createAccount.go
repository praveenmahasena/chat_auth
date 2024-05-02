package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/postgres"
)

func CreateAccount(gctx *gin.Context) {
	var usr = postgres.NewUser()

	if err := gctx.BindJSON(&usr); err != nil {
		gctx.JSONP(http.StatusNotAcceptable, "")
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), time.Duration(10)*time.Minute)
	defer cancel()
	if exist := usr.Check(ctx); exist {
		return
	}
	gctx.JSONP(http.StatusOK, "done")
}
