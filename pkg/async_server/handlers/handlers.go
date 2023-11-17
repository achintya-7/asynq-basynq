package handlers

import (
	"asqnq-test/pkg/models"
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

func HandleWelcomeEmailTask(ctx context.Context, t *asynq.Task) error {
	var req models.WelcomeEmailTaskModel
	if err := json.Unmarshal(t.Payload(), &req); err != nil {
		return err
	}

	log.Printf(" [*] Send Welcome Email to User %s with message %s", req.Email, req.Message)
    return nil
}

func HandleTrainModelTask(ctx context.Context, t *asynq.Task) error {
	var req models.TrainModelTaskModel
	if err := json.Unmarshal(t.Payload(), &req); err != nil {
		return err
	}

	log.Printf(" [*] Train Model %s with GPU %s", req.ModelID, req.GPU)
	return nil
}
