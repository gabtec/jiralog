package utils

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"time"
)

func InitNewWeek(taskId int, targetPath string) {
	t := time.Now()
	_, week := t.ISOWeek()

	dirName := fmt.Sprintf("week-%d", week)
	fullTargetPath := path.Join(targetPath, dirName)
	
	err := os.MkdirAll(fullTargetPath, 0755)
    if err != nil {
        fmt.Println("Error creating directory:", err)
        return
    }
	fmt.Printf("[INFO]: Directory '%s' created successfully!", fullTargetPath)

	filePath := path.Join(fullTargetPath, "worklog.yaml")
	buildBaseFileFromTemplate(filePath)
}

func buildBaseFileFromTemplate(filePath string) {
    // Define a YAML template string with placeholders using Go template syntax
    const yamlTemplate = `
data:
  "{{ .StartDate }}": # SEGUNDA
    {{ .Prefix }}-{{ .TaskID }}:
      description: Daily
      timeSpent: "30m"
`

    // Define the data to fill in the template
    data := struct {
        StartDate  string
        TaskID int
		Prefix string
    }{
        StartDate:  "2025-11-24",
        TaskID: 8446,
		Prefix: GetStringEnv("JIRA_PREFIX", "ABC"),
    }

    // Parse the template string
    tmpl, err := template.New("yaml").Parse(yamlTemplate)
    if err != nil {
        panic(err)
    }

    // Create or open the output file
    file, err := os.Create(filePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Execute the template with data and write output to the file
    err = tmpl.Execute(file, data)
    if err != nil {
        panic(err)
    }
}
