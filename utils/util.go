package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	pulseURL = "https://application-9f.ubi6nzyvztg.us-south.codeengine.appdomain.cloud"
)

func GetPlanFile(planFilePath string) ([]byte, error) {
	// Open plan jsonFile
	planFile, err := os.Open(planFilePath)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to open plan file: %s", err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer planFile.Close()

	planBytes, err := ioutil.ReadAll(planFile)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to read plan file: %s", err)
	}
	return planBytes, nil
}
func GetPulseUrl() string {
	if os.Getenv("PULSE_URL") != "" {
		return os.Getenv("PULSE_URL")
	}
	return pulseURL
}
