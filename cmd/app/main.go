package main

import (
	"asqnq-test/pkg/api"
	asyncserver "asqnq-test/pkg/async_server"
	"log"
	"time"
)

func main() {
	log.Printf("Starting %d async server...", asyncserver.Servers)
	asyncserver.StartAsynqServers()

	// Wait for async servers to start
	time.Sleep(1 * time.Second)

	log.Println("Starting API server...")
	server := api.NewServer()
	server.Run()
}
