package asyncserver

import (
	"asqnq-test/pkg/async_server/handlers"
	"asqnq-test/pkg/asynq_client/tasks"
	"log"

	"github.com/hibiken/asynq"
)

const (
	Servers  = 3
	redisAddr = "127.0.0.1:6379"
)

func StartAsynqServers() {
	for i := 0; i < Servers; i++ {
		go startAsynqServer()
	}
}

func startAsynqServer() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.WelcomeEmailTask, handlers.HandleWelcomeEmailTask)
	mux.HandleFunc(tasks.TrainModelTask, handlers.HandleTrainModelTask)

	if err := srv.Run(mux); err != nil {
		log.Println("Error running async server")
		log.Fatal(err)
	}
}
