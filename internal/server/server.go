package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/chat-app/internal/helpers"
	"github.com/praveenmahasena647/chat-app/internal/postgres"
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

	g.GET("/profile", handlers.RouteGurd, handlers.Profile)
	g.GET("/verify", handlers.RouteGurd, handlers.Verify)

	g.GET("/mark/:id", handlers.MarkVerify)

	g.POST("/login", func(ctx *gin.Context) {
		u := postgres.NewLoginCredentials()
		if err := ctx.Bind(u); err != nil {
			ctx.JSONP(http.StatusMethodNotAllowed, "the credentials are not allowed")
			return
		}
		usrInfo, usrInfoErr := u.Login(ctx)
		if usrInfoErr != nil {
			ctx.JSONP(http.StatusNotFound, usrInfoErr)
			return
		}
		if err := helpers.VerifyPassword(usrInfo.Password, u.Password); err != nil {
			ctx.JSONP(http.StatusNotFound, err)
			return
		}
		JWT, JWTErr := helpers.GenerateJWT(u.EmailID)
		if JWTErr != nil {
			ctx.JSONP(http.StatusNotFound, JWTErr)
			return
		}
		ctx.JSONP(http.StatusOK, JWT)
	})

	g.GET("/chatroom", handlers.RouteGurd, func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	return g.Run(s.ListenAddr)
}
