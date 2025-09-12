package utils

import (
	"fmt"
	"gabtec/log-hours/models"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func BuildTable(sd models.SData) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Dia", "Task", "Description", "Start at", "Time Spent"})

	var lastDate string
	var durations []string

	fmt.Println("[INFO]: Results table:")
	for i, entry := range sd {
		if i > 0 && entry.Date != lastDate {
			total := CalculateTotalHoursOfDay(durations)
			timeRow := table.Row{"", "", "", "--> Total", total}
			t.AppendRow(timeRow)
			t.AppendSeparator()
			durations = durations[:0] // reset slice
		}
		t.AppendRow(table.Row{entry.Date, entry.TaskID, entry.Description, entry.Start, entry.TimeSpent})
		durations = append(durations, entry.TimeSpent) // <-- move here
		lastDate = entry.Date
	}

	// Print total for the last date group
	if len(durations) > 0 {
		total := CalculateTotalHoursOfDay(durations)
		timeRow := table.Row{"", "", "", "--> Total", total}
		t.AppendRow(timeRow)
	}

	t.Render()
}

func printHeader(idx int) string {
	switch idx {
	case 0:
		return "MONDAY    - SEGUNDA: "
	case 1:
		return "TUESDAY   - TERÇA  : "
	case 2:
		return "WEDNESDAY - QUARTA : "
	case 3:
		return "THURSDAY  - QUINTA : "
	case 4:
		return "FRIDAY    - SEXTA  : "
	}
	return ""
}

func BuilReport(sd models.SData, prefix, tag string) {

	fullTag := fmt.Sprintf("%s - %s", prefix, tag)

	var lastDate string
	var day = 0

	fmt.Println("[INFO]: Printing Report...")

	for _, entry := range sd {

		if entry.Date != lastDate {
			fmt.Println("")
			fmt.Println(printHeader(day), entry.Date)
			fmt.Println("---------------------------------")
			fmt.Println(fullTag)
			day += 1
		}

		row := entry.TaskID + ": " + entry.Description + " (" + entry.TimeSpent + ")"
		fmt.Println(row)
		lastDate = entry.Date

	}
}
