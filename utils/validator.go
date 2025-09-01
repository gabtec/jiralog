package utils

import (
	"fmt"
	"gabtec/log-hours/models"
	"log"
	"os"
	"regexp"
)

func MustValidate(e models.Entry) bool {
	if !isValidDate(e.Date) {
		log.Fatal("Date is not in a valid format, like YYYY-MM-DD \nYou have: ", e.Date)
	}
	if !isValidTime(e.Start) {
		log.Fatal("Start time is not in a valid format, like HH:MM \nYou have: ", e.Start)
	}
	if !isValidDuration(e.TimeSpent) {
		log.Fatal("Time spent is not in a valid format, use only [1-8]h, combined or not with 30m. \nYou have: ", e.TimeSpent)
	}

	// if date not from current week...
	taskPrefix := os.Getenv("JIRA_PREFIX")

	if taskPrefix == "" {
		log.Fatal("Missing one required environment variables: JIRA_PREFIX")
	}

	if !isValidTicketID(e.TaskID, taskPrefix) {
		msg := fmt.Sprintf("Task ID is not in a valid format, like %s-1234 \nYou have: %s", taskPrefix, e.TaskID)
		log.Fatal(msg)
	}

	return true
}

func isValidDate(date string) bool {
	pattern := `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(date)
}

func isValidTicketID(ticket, prefix string) bool {
	pattern := fmt.Sprintf("^%s-\\d+$", prefix)
	matched, _ := regexp.MatchString(pattern, ticket)
	return matched
}

func isValidTime(timeStr string) bool {
	// Regex to match HH:MM where HH is 00-23 and MM is 00-59
	pattern := `^(?:[01][0-9]|2[0-3]):[0-5][0-9]$`
	matched, _ := regexp.MatchString(pattern, timeStr)
	return matched
}

func isValidDuration(d string) bool {
	// Regex matches exactly one of the allowed strings
	pattern := `^([1-8]+h)?\s?(30m)?$`
	matched, _ := regexp.MatchString(pattern, d)
	return matched
}
