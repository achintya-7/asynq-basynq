package main

import (
	"asqnq-test/pkg/api"
	asyncserver "asqnq-test/pkg/async_server"
	"log"
)

func main() {
	log.Println("Starting asynq server...")
	go asyncserver.StartAsynqServer()

	log.Println("Starting API server...")
	server := api.NewServer()
	server.Run()
}
