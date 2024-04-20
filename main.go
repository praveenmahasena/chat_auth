package main

import (
	"log"

	"github.com/praveenmahasena647/chat-app/cmd"
)

func main() {
	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}
}
