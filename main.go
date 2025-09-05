package main

import (
	"fmt"
	"gabtec/log-hours/constants"
	"gabtec/log-hours/models"
	"gabtec/log-hours/utils"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type WorkLog struct {
	Data map[string]interface{} `yaml:"data"`
}

func main() {
	var dryRun = false
	var report = false

	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("[ERROR]: Only one argument is allowed.")
		fmt.Println("")
		utils.ShowUsage()
		os.Exit(1)
	}
	for _, arg := range args {
		if arg == "-v" || arg == "--version" || arg == "version" {
			fmt.Printf("%s (r), version: %s\n", constants.AppName, constants.Version)
			os.Exit(0)
		}

		if arg == "-d" || arg == "--dry-run" {
			dryRun = true
			fmt.Println("[INFO]: Running in dry mode")
		}

		if arg == "-r" || arg == "--report" {
			report = true
			fmt.Println("[INFO]: Running in report mode")
		}

		if !(dryRun || report) {
			fmt.Println("[ERROR]: Unknown flag.")
			fmt.Println("")
			utils.ShowUsage()
			os.Exit(2)
		}
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read the YAML file
	fileContent, err := os.ReadFile("worklog.yaml")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	var w WorkLog
	var sd = make(models.SData, 0)

	// Unmarshal YAML into the struct
	err = yaml.Unmarshal(fileContent, &w)
	if err != nil {
		log.Fatalf("error unmarshaling YAML: %v", err)
	}

	// load
	// Access nested keys step-by-step with type assertions
	data := w.Data
	if data == nil {
		log.Fatal("data key not found or not a map")
	}

	for date, tasksRaw := range data {
		var entry = models.Entry{}

		entry.Date = date

		tasks, ok := tasksRaw.(map[string]interface{})
		if !ok {
			log.Fatal("tasks is not a map")
		}

		for taskID, detailsRaw := range tasks {
			entry.TaskID = taskID
			details, ok := detailsRaw.(map[string]interface{})
			if !ok {
				log.Fatal("details is not a map")
			}

			start, _ := details["start"].(string)

			if start == "" {
				start = models.DefaultStart
			}

			timeSpent, _ := details["timeSpent"].(string)
			desc, _ := details["description"].(string)

			entry.Start = start
			entry.TimeSpent = timeSpent
			entry.Description = desc
			// fmt.Printf("  Task %s: start: %s, timeSpent: %s\n", taskID, start, timeSpent)

			utils.MustValidate(entry)

			// Append a copy of entry for each task
			sd = append(sd, entry)
		}

	}

	// baseUrl := GetStringEnv("JIRA_BASE_URL", "default")
	// apiToken := GetStringEnv("JIRA_API_TOKEN", "default")
	baseUrl := os.Getenv("JIRA_BASE_URL")
	apiToken := os.Getenv("JIRA_API_TOKEN")

	if baseUrl == "" || apiToken == "" {
		log.Fatal("Missing one or more required environment variables: JIRA_BASE_URL, JIRA_API_TOKEN")
	}

	// make api calls
	var wg sync.WaitGroup
	for _, record := range sd {
		if !(dryRun || report) {
			wg.Add(1)
			go func(rec models.Entry) {
				defer wg.Done()
				utils.UploadWorkLog(rec, baseUrl, apiToken)
			}(record)
		}
	}
	wg.Wait()

	// table
	sortedData := utils.SortTableData(sd)

	if dryRun {
		utils.BuildTable(sortedData)
	}

	if report {
		utils.BuilKantataReport(sortedData)
	}

}
