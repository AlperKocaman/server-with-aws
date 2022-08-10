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

	err := config.InitializeConfigForApp()
	if err != nil {
		log.Fatal("error while reading config, exiting.")
	}

	log.Println("Starting http server...")
	err = server.InitAndRunServer()
	if err != nil {
		log.Fatal("error while starting server, exiting.")
		return
	}
}
