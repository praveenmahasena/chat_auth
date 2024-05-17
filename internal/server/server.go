package server

import (
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
	var g = gin.Default()

	g.Use(helpers.CORS())

	g.POST("/create", handlers.CreateAccount)
	g.Use(helpers.DecodeToken())
	g.GET("/profile", handlers.Profile)
	g.GET("/verify", handlers.Verify)

	return g.Run(s.ListenAddr)
}
