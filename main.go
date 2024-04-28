package main

import (
	"log"

	"github.com/praveenmahasena647/chat-app/internal"
)

func main() {
	if err := internal.Start(); err != nil {
		log.Fatalln(err)
	}
}
