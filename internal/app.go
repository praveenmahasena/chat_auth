package internal

import (
	"github.com/praveenmahasena647/chat-app/internal/postgres"
	"github.com/praveenmahasena647/chat-app/internal/server"
)

func Start() error {
	DBConnectionErr := postgres.Connect()
	if DBConnectionErr != nil {
		return DBConnectionErr
	}
	s := server.New(":42069")
	return s.Run()
}
