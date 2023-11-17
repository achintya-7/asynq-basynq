package models

type WelcomeEmailTaskModel struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

type TrainModelTaskModel struct {
	ModelID string `json:"model_id" binding:"required"`
	CPU     string `json:"cpu" binding:"required"`
	GPU     string `json:"gpu" binding:"required"`
	Memory  int16  `json:"memory" binding:"required"`
}

type ModelTrainedEmailTaskModel struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}
