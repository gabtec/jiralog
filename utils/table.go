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
