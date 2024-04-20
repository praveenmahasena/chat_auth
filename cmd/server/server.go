package server

import (
	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/cmd/helpers"
	"github.com/praveenmahasena647/chat-app/cmd/server/handlers"
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

	return g.Run(s.ListenAddr)
}
