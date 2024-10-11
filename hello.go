package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	request := GymTrainingPlanRequest{
		FitnessLevel:           1,
		TrainingGoal:           1,
		TimeCommitmentInMinute: 60,
		AdditionalContext:      "Make it all legs",
	}

	trainingPlan, trainingPlanError := GetGymTrainingPlan(request)
	if trainingPlanError != nil {
		fmt.Println("Error getting training plan:", trainingPlanError)
		return
	}

	fmt.Println(trainingPlan)
}

type GymTrainingPlanRequest struct {
	FitnessLevel           int    `json:"fitnessLevel"`
	TrainingGoal           int    `json:"trainingGoal"`
	TimeCommitmentInMinute int    `json:"timeCommitmentInMinute"`
	AdditionalContext      string `json:"additionalContext"`
}

func GetGymTrainingPlan(request GymTrainingPlanRequest) (string, error) {
	requestJson, rerequestJsonError := json.Marshal(request)
	if rerequestJsonError != nil {
		fmt.Println("Error marshalling request to JSON:", rerequestJsonError)
		return "", rerequestJsonError
	}

	response, responseError := http.Post("https://playground-core.onrender.com/llmrecommendation/gym-training-plan", "application/json", bytes.NewBuffer(requestJson))
	if responseError != nil {
		fmt.Println("Error sending request:", responseError)
		return "", responseError
	}
	defer response.Body.Close()

	body, bodyError := io.ReadAll(response.Body)
	if bodyError != nil {
		fmt.Println("Error reading response body:", bodyError)
		return "", bodyError
	}

	return string(body), nil
}
