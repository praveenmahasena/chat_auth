package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/postgres"
)

func Profile(gctx *gin.Context) {
	var email, _ = gctx.Get("Email")
	var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()
	var verified, err = postgres.IsVerified(ctx, email.(string))
	if err != nil {
		gctx.JSONP(http.StatusInternalServerError, "")
		return
	}
	if !verified {
		gctx.JSONP(http.StatusOK, false)
		return
	}
	gctx.JSONP(http.StatusOK, true)
}
