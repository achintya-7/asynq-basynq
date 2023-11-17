package tasks

import (
	"asqnq-test/pkg/models"
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	WelcomeEmailTask      = "welcome_email"
	TrainModelTask        = "train_model"
	ModelTrainedEmailTask = "model_trained_email"
)

func NewWelcomeEmailTask(email string) (*asynq.Task, error) {
	emailPayload := models.WelcomeEmailTaskModel{
		Email:   email,
		Message: "Welcome to our service!",
	}

	payload, err := json.Marshal(emailPayload)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(WelcomeEmailTask, payload), nil
}

func NewTrainModelTask(model models.TrainModelTaskModel) (*asynq.Task, error) {
	payload, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TrainModelTask, payload), nil
}

func NewModelTrainedEmailTask(email string) (*asynq.Task, error) {
	emailPayload := models.ModelTrainedEmailTaskModel{
		Email:   email,
		Message: "Your model is trained and ready to use!",
	}

	payload, err := json.Marshal(emailPayload)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(ModelTrainedEmailTask, payload), nil
}
