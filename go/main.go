package main

import (
	"github.com/AlperKocaman/server-with-aws/cmd/config"
	"github.com/AlperKocaman/server-with-aws/core/server"
	"log"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Fatal("panic occurred in main")
		}
	}()

	err := config.InitializeConfig()
	if err != nil {
		log.Fatal("error while reading config, exiting.")
	}

	log.Println("Starting http server...")
	err = server.InitializeServer()
	if err != nil {
		log.Fatal("error while starting server, exiting.")
		return
	}
}
