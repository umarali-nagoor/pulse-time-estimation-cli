package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/IBM-Cloud/pulse-time-estimation-cli/utils"
)

func DeletePredictTime(ID string) (map[string]interface{}, error) {

	var deleteResponse map[string]interface{}

	req, err := http.NewRequest("DELETE", utils.GetPulseUrl()+"/api/v1/predictor/"+ID, nil)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] DELETE new request failure:  %s", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error deleting JobID:  %s", err)
	}
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("[ERROR] error deleting jobID.\nResponse: %+v\nError:%s", resp, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] unable to read response body:  %s", err)
	}
	unmarshallErr := json.Unmarshal([]byte(body), &deleteResponse)
	if unmarshallErr != nil {
		return nil, fmt.Errorf("[ERROR] unmarshall error: %+v", unmarshallErr)
	}
	return deleteResponse, nil
}
