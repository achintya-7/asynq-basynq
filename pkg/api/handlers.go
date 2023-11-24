package api

import (
	"asqnq-test/pkg/asynq_client/tasks"
	"asqnq-test/pkg/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) welcomeHandler(c *gin.Context) {
	var req models.WelcomeApiModel
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	t, err := tasks.NewWelcomeEmailTask(req.Email)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	info, err := s.asynqClient.Enqueue(t)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, info.ID)
}

func (s *Server) trainHandler(c *gin.Context) {
	var req models.TrainModelTaskModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	t, err := tasks.NewTrainModelTask(req)		
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	info, err := s.asynqClient.Enqueue(t)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, info.ID)
}

func (s *Server) modelTrainedHandler(c *gin.Context) {
	var req models.WelcomeApiModel
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	t, err := tasks.NewModelTrainedEmailTask(req.Email)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	info, err := s.asynqClient.Enqueue(t)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, info.ID)
}
