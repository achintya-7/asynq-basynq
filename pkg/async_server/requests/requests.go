package requests

import (
	"asqnq-test/pkg/models"
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	BASE_URL = "http://localhost:2205"
)


func WelcomeRequest(email string) error {
	url := BASE_URL + "/welcome?email=" + email
	
	_, err := http.Get(url)
	if err != nil {
		return err
	}

	return nil
}

func TrainRequest(trainingData models.TrainModelTaskModel) error {
	url := BASE_URL + "/train"

	body, err := json.Marshal(trainingData)
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	return nil
}

func TrainedRequest(email string) error {
	url := BASE_URL + "/trained?email=" + email
	
	_, err := http.Get(url)
	if err != nil {
		return err
	}

	return nil
}