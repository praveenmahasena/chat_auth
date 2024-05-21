package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/postgres"
)

func Profile(gctx *gin.Context) {
	emailID, _ := gctx.Get("Email")
	verified, verifyErr := postgres.IsVerified(gctx, emailID.(string))
	if verifyErr != nil {
		gctx.JSONP(http.StatusInternalServerError, "")
		return
	}
	if !verified {
		gctx.JSONP(http.StatusOK, false)
		return
	}
	gctx.JSONP(http.StatusOK, true)
}
