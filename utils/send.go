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

func UploadWorkLog(e models.Entry, baseUrl, apiToken string) {

	apiVersion := "2" // "v3 is for Jira Cloud - not self-hosted"
	// baseUrl := GetStringEnv("JIRA_BASE_URL", "default")
	// apiToken := GetStringEnv("JIRA_API_TOKEN", "default")
	// baseUrl := os.Getenv("JIRA_BASE_URL")
	// apiToken := os.Getenv("JIRA_API_TOKEN")

	// if baseUrl == "" || apiToken == "" {
	// 	log.Fatal("Missing one or more required environment variables: JIRA_BASE_URL, JIRA_API_TOKEN")
	// }

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

	if resp.StatusCode != 201 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Fatalf("Non-OK response status: %d\nBody: %s", resp.StatusCode, string(bodyBytes))
	} else {
		fmt.Println("[ OK ]: Success")
		fmt.Printf("Response: %v\n", resp.Body)
	}

	// make http call

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("Failed to read response body: %v", err)
	// }

	// var result interface{}
	// if err := json.Unmarshal(body, &result); err != nil {
	// 	log.Fatalf("Failed to parse JSON response: %v", err)
	// }

	// resultJson, err := json.MarshalIndent(result, "", "  ")
	// if err != nil {
	// 	log.Fatalf("Failed to format JSON: %v", err)
	// }

	// fmt.Println("Get Bearer Token response:")
	// fmt.Println(string(resultJson))
}
