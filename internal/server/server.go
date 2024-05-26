package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/helpers"
	"github.com/praveenmahasena647/chat-app/internal/server/handlers"
)

type Server struct {
	ListenAddr string
}

func New(port string) *Server {
	return &Server{
		ListenAddr: port,
	}
}

func (s *Server) Run() error {
	g := gin.Default()

	g.Use(helpers.CORS())

	g.POST("/create", handlers.CreateAccount)

	g.GET("/profile", handlers.RouteGuard, handlers.Profile)
	g.GET("/verify", handlers.RouteGuard, handlers.Verify)

	g.GET("/mark/:id", handlers.MarkVerify)

	g.POST("/login", handlers.Login)

	g.GET("/chatroom", handlers.RouteGuard, func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	return g.Run(s.ListenAddr)
}
