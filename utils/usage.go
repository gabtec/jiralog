package utils

import (
	"fmt"
	"gabtec/log-hours/constants"
)

func ShowUsage() {

	fmt.Printf("usage: %s [flag]\n", constants.AppName)
	fmt.Println("flags:")
	fmt.Println(" -v, --version, version -- show version")
	fmt.Println(" -d, --dry-run          -- run in dry mode (no http calls will be made)")
	fmt.Println(" -r, --report           -- run in report mode (no http calls will be made)")
	fmt.Println("data file:")
	fmt.Println("  by default it read data from a file named: 'worklog.yaml' \n")
}
