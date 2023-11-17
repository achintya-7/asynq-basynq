package api

import (
	asynqclient "asqnq-test/pkg/asynq_client"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type Server struct {
	asynqClient *asynq.Client
	router      *gin.Engine
}

func NewServer() *Server {
	client := asynqclient.NewClient()

	server := &Server{
		asynqClient: client,
	}

	server.setup()

	return server
}

func (s *Server) setup() {
	// New so as to not modify the default gin router
	// It will not give out logs so we can see the logs from the asynq server
	r := gin.New()

	r.GET("/welcome", s.welcomeHandler)
	r.POST("/train", s.trainHandler)

	s.router = r
}

func (s *Server) Run() {
	s.router.Run("127.0.0.1:2205")
}
