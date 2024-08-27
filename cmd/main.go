package main

import (
	"log"

	"github.com/faraz-wq/video-call-app/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
