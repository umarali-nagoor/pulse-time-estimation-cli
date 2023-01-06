package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM-Cloud/pulse-time-estimation-cli/api"
)

func GetPredictionStatus(JobId string) error {
	log.Printf("[INFO] ------- Getting prediction status for JobId %s -------", JobId)
	predictedstatusResponse, predictedstatusResponseError := api.GetPredictedJobStatus(JobId)
	if predictedstatusResponseError != nil {
		return fmt.Errorf("%+v", predictedstatusResponseError)
	}
	predictedstatusResponseBytes, err := json.Marshal(predictedstatusResponse)
	if err != nil {
		return fmt.Errorf("[ERROR] marshall error: %+v", err)
	}
	fmt.Printf("%+v\n", string(predictedstatusResponseBytes[:]))
	log.Printf("[INFO] ------- Got prediction status for JobID - %s -------", JobId)
	return nil
}
