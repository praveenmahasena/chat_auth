package internal

import (
	"github.com/praveenmahasena647/chat-app/internal/postgres"
	"github.com/praveenmahasena647/chat-app/internal/server"
)

func Start() error {
	if err := postgres.Connect(); err != nil {
		return err
	}
	var s = server.New(":42069")
	return s.Run()
}
