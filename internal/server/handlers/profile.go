package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/postgres"
)

func Profile(gctx *gin.Context) {
	var email, _ = gctx.Get("Email")
	var verified, err = postgres.IsVerified(gctx, email.(string))
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
