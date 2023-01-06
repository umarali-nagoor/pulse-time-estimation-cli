package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/IBM-Cloud/pulse-time-estimation-cli/utils"
)

func PredictTime(request []byte) (map[string]interface{}, error) {
	var predictResponse map[string]interface{}

	predictRequest := bytes.NewBuffer(request)

	resp, err := http.Post(utils.GetPulseUrl()+"/api/v1/predictor", "application/json", predictRequest)
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("[ERROR] error predicting time for given plan.\nResponse: %+v\nError:%s", resp, err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] unable to read response body:  %s", err)
	}
	json.Unmarshal([]byte(body), &predictResponse)
	return predictResponse, nil
}
