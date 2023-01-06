package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM-Cloud/pulse-time-estimation-cli/api"
)

func DeletePrediction(JobId string) error {
	log.Printf("[INFO] ------- Deleting prediction for JobId %s -------", JobId)
	deleteResponse, deleteResponseResponseError := api.DeletePredictTime(JobId)
	if deleteResponseResponseError != nil {
		return fmt.Errorf("%+v", deleteResponseResponseError)
	}
	deleteResponseBytes, err := json.Marshal(deleteResponse)
	if err != nil {
		return fmt.Errorf("[ERROR] marshall error: %+v", err)
	}
	fmt.Printf("%+v\n", string(deleteResponseBytes[:]))
	log.Printf("[INFO] ------- Deleted JobID - %s -------", JobId)
	return nil
}
