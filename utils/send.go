package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gabtec/log-hours/models"
	"io"
	"log"
	"net/http"
)

type JiraResponse struct {
	IssueId   string `json:"issueId"`
	LogId     string `json:"id"`
	Created   string `json:"created"`
	TimeSpent string `json:"timeSpent"`
	Self      string `json:"self"`
}

func UploadWorkLog(e models.Entry, baseUrl, apiToken string) {

	apiVersion := "2" // "v3 is for Jira Cloud - not self-hosted"

	compositeUrl := fmt.Sprintf("%s/rest/api/%s/issue/%s/worklog", baseUrl, apiVersion, e.TaskID)

	body := map[string]string{
		"started":   fmt.Sprintf("%sT%s:00.000+0000", e.Date, e.Start),
		"timeSpent": e.TimeSpent,
		"comment":   e.Description, // in v3 comment is a more complex object
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Failed to marshal body: %v", err)
	}
	postRequest, err := http.NewRequest("POST", compositeUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	postRequest.Header.Set("Authorization", "Bearer "+apiToken)
	postRequest.Header.Set("Accept", "application/json")
	postRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(postRequest)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 201 {
		log.Fatalf("Non-OK response status: %d\nBody: %s", resp.StatusCode, string(bodyBytes))
	} else {
		fmt.Println("[ OK ]: Success")
		var responseObj JiraResponse
		err := json.Unmarshal(bodyBytes, &responseObj)
		if err != nil {
			fmt.Println("[ERROR]: unable to read response")
		} else {
			fmt.Printf("Logged %s at %s \n", responseObj.TimeSpent, responseObj.Self)
		}
	}
}
