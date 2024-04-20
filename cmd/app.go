package cmd

import (
	"github.com/praveenmahasena647/chat-app/cmd/postgres"
	"github.com/praveenmahasena647/chat-app/cmd/server"
)

func Start() error {
	if err := postgres.Connect(); err != nil {
		return err
	}
	var s = server.New(":42069")
	return s.Run()
}
