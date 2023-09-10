package tradervue

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Import Exectutions to Tradervue as per:
// https://github.com/tradervue/api-docs/blob/master/imports.md
func Import(executions []Execution, username string, password string, tags []string, accountTag *string) error {

	ac := ""
	if accountTag != nil {
		ac = *accountTag
	}

	// build import request
	importRequest := ImportRequest{
		AllowDuplicates:   false,
		OverlayCommisions: false,
		Tags:              tags,
		AccountTag:        ac,
		Executions:        executions,
	}

	// encode request to JSON
	body, err := json.Marshal(importRequest)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to encode imprt request %s", err.Error()))
	}

	// create POST request
	client := &http.Client{}
	url := "https://www.tradervue.com/api/v1/imports"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to create request %s", err.Error()))
	}

	// set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Sierrachart2Tradervue")
	req.Header.Set("Content-Type", "application/json")

	// set auth
	req.SetBasicAuth(username, password)

	httpResponse, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to POST request %s", err.Error()))
	}

	var response ImportResponse
	defer httpResponse.Body.Close()
	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to decode POST request response %s", err.Error()))
	}

	// check response
	if httpResponse.StatusCode != http.StatusOK || response.Status != "queued" {
		return errors.New(response.Status)
	}

	if response.Status == "queued" {
		log.Println("Success! Tradervue import request was not queued. Check your Tradervue account now.")
		log.Println("https://apteros.tradervue.com/trades")
	} else {
		log.Println(fmt.Sprintf("Tradervue import request was submitted successfully, but we got an unexpected status response: %s\n", response.Status))
		log.Println("Check your Tradervue account to make sure trades imported as expectd.")
		log.Println("https://apteros.tradervue.com/trades")
	}

	return nil
}
