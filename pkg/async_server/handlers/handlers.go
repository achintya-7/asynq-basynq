package handlers

import (
	"asqnq-test/pkg/async_server/requests"
	"asqnq-test/pkg/models"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func HandleWelcomeEmailTask(ctx context.Context, t *asynq.Task) error {
	var req models.WelcomeEmailTaskModel
	if err := json.Unmarshal(t.Payload(), &req); err != nil {
		return err
	}

	time.Sleep(time.Second)

	trainingData := models.TrainModelTaskModel{
		ModelID: "1",
		GPU:     "NVIDIA 1080ti",
		CPU:     "Intel i7 8700k",
		Memory:  16,
	}

	err := requests.TrainRequest(trainingData)
	if err != nil {
		return err
	}

	log.Printf(" [*] Sent Welcome Email to User %s with message %s", req.Email, req.Message)
	return nil
}

func HandleTrainModelTask(ctx context.Context, t *asynq.Task) error {
	var req models.TrainModelTaskModel
	if err := json.Unmarshal(t.Payload(), &req); err != nil {
		return err
	}

	time.Sleep(time.Second * 3)

	emailID := "example@email.com"
	err := requests.TrainedRequest(emailID)
	if err != nil {
		return err
	}

	log.Printf(" [*] Train Model %s with GPU %s", req.ModelID, req.GPU)
	return nil
}

func HandleModelTrainedEmailTask(ctx context.Context, t *asynq.Task) error {
	var req models.ModelTrainedEmailTaskModel
	if err := json.Unmarshal(t.Payload(), &req); err != nil {
		return err
	}

	time.Sleep(time.Second)

	log.Printf(" [*] Sent Model Trained Email to User %s with message %s", req.Email, req.Message)
	return nil
}
