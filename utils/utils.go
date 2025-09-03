package utils

import (
	"gabtec/log-hours/models"
	"os"
	"sort"
)

func GetStringEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

func SortTableData(sd models.SData) models.SData {
	sort.Slice(sd, func(i, j int) bool {
		return sd[i].Date < sd[j].Date
	})

	return sd
}
