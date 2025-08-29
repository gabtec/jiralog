package models

const (
	DefaultStart = "09:00"
)

type Entry struct {
	Date        string `yaml:"date"`
	TaskID      string `yaml:"taskId"`
	Start       string `yaml:"start"`
	TimeSpent   string `yaml:"timeSpent"`
	Description string `yaml:"description,omitempty"`
}

// SData = structured data
type SData []Entry
