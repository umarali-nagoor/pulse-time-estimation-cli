package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/IBM-Cloud/pulse-time-estimation-cli/utils"
)

func GetPredictedJobStatus(ID string) (map[string]interface{}, error) {

	var predictedResponse map[string]interface{}

	resp, err := http.Get(utils.GetPulseUrl() + "/api/v1/predictor/" + ID + "/status")
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("[ERROR] error executing GetPredictedJobStatus.\nResponse: %+v\nError:%s", resp, err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] unable to read response body:  %s", err)
	}
	unmarshallErr := json.Unmarshal([]byte(body), &predictedResponse)
	if unmarshallErr != nil {
		return nil, fmt.Errorf("[ERROR] unmarshall error: %+v", unmarshallErr)
	}
	return predictedResponse, nil
}
