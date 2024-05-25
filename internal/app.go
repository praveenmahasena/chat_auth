package internal

import (
	"github.com/praveenmahasena647/chat-app/internal/postgres"
	"github.com/praveenmahasena647/chat-app/internal/server"
	"github.com/spf13/viper"
)

func Start() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	DBConnectionErr := postgres.Connect()
	if DBConnectionErr != nil {
		return DBConnectionErr
	}

	s := server.New(":42069")
	return s.Run()
}
