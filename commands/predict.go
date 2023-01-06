package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM-Cloud/pulse-time-estimation-cli/api"
	"github.com/IBM-Cloud/pulse-time-estimation-cli/payloads"
	"github.com/IBM-Cloud/pulse-time-estimation-cli/utils"
)

func Predict(planFilePath string) error {
	log.Printf("[INFO] ------- Executing predict for given plan file %s -------", planFilePath)

	predictRequestPayload, err := utils.GetPlanFile(planFilePath)
	if err != nil {
		return fmt.Errorf("[ERROR] payload error: %s", err)
	}

	predictTimeResponse, predictErr := api.PredictTime(predictRequestPayload)
	if predictErr != nil {
		return fmt.Errorf("%s", predictErr)
	}
	if predictTimeResponse == nil && predictTimeResponse["jobID"] == nil {
		return fmt.Errorf("[ERROR] predictTime response error: response is nil")
	}
	predictedJobIDString := fmt.Sprintf("%+v", predictTimeResponse["jobID"].(float64))

	TimeEstimationResponse := payloads.TimeEstimationResult{}
	for {
		predictedResponse, predictedResponseError := api.GetPredictTime(predictedJobIDString)
		if predictedResponseError != nil {
			return fmt.Errorf("%+v", predictedResponseError)
		}
		predictedResponseBytes, marshallErr := json.Marshal(predictedResponse)
		if marshallErr != nil {
			return fmt.Errorf("[ERROR] marshall error: %+v", marshallErr)
		}
		unmarshallErr := json.Unmarshal(predictedResponseBytes, &TimeEstimationResponse)
		if unmarshallErr != nil {
			return fmt.Errorf("[ERROR] unmarshall error: %+v", unmarshallErr)
		}

		if len(TimeEstimationResponse.Resources) != 0 {
			fmt.Printf("%+v\n", string(predictedResponseBytes[:]))
			break
		}
		log.Println("[INFO] Waiting for prediction to complete")
	}
	log.Printf("[INFO] ------- Prediction completed for given plan file %s -------", planFilePath)
	return nil
}
