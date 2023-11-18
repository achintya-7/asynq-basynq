package main

import (
	"asqnq-test/pkg/api"
	asyncserver "asqnq-test/pkg/async_server"
	"log"
)

func main() {
	log.Printf("Starting %d async server...", asyncserver.Servers)
	asyncserver.StartAsynqServers()

	log.Println("Starting API server...")
	server := api.NewServer()
	server.Run()
}
