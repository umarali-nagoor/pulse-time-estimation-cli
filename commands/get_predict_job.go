package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM-Cloud/pulse-time-estimation-cli/api"
)

func GetPrediction(JobId string) error {
	log.Printf("[INFO] ------- Getting prediction for JobId %s -------", JobId)
	predictedResponse, predictedResponseError := api.GetPredictTime(JobId)
	if predictedResponseError != nil {
		return fmt.Errorf("%+v", predictedResponseError)
	}
	predictedResponseBytes, err := json.Marshal(predictedResponse)
	if err != nil {
		return fmt.Errorf("[ERROR] marshall error: %+v", err)
	}
	fmt.Printf("%+v\n", string(predictedResponseBytes[:]))
	log.Printf("[INFO] ------- Got prediction for JobID - %s -------", JobId)
	return nil
}
