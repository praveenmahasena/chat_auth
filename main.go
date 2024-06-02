package main

import (
	"log"

	"github.com/praveenmahasena647/chat-app/internal"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	if err := internal.Start(); err != nil {
		log.Fatalln(err)
	}
}
