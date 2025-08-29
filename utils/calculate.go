package utils

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func CalculateTotalHoursOfDay(durations []string) string {
	var totalDuration time.Duration
	splitedTime := splitCompositeTimes(durations)

	for _, d := range splitedTime {
		dur, err := time.ParseDuration(d)
		if err != nil {
			fmt.Printf("Error parsing duration %s: %v\n", d, err)
			return "n/a"
		}
		totalDuration += dur
	}

	max, err := time.ParseDuration("8h")
	if err != nil {
		log.Fatal("Failed to parse max duration: ", err)
	}

	if totalDuration > max {
		log.Fatal("Your total working hours, per day, can not be > 8h.\nYou have: ", totalDuration)
	}
	// fmt.Printf("Total duration: %v\n", totalDuration)
	return fmt.Sprintf("%v", showOnlyHoursIfRestIsZero(totalDuration))
}

func splitCompositeTimes(durations []string) []string {
	var aux []string
	for _, d := range durations {
		parts := strings.Split(d, " ")
		aux = append(aux, parts...)
	}
	fmt.Println("debug composite")
	fmt.Printf("%v\n", aux)
	return aux
}

func showOnlyHoursIfRestIsZero(total time.Duration) string {
	hours := int(total.Hours())
	minutes := int(total.Minutes()) % 60
	seconds := int(total.Seconds()) % 60

	if minutes == 0 && seconds == 0 {
		return fmt.Sprintf("%dh", hours)
	} else {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
}
