package io

type Task struct {
	Name         string   `yaml:"name"`
	ScheduleType string   `yaml:"schedule_type"`
	Run          string   `yaml:"run"`
	Arguments    []string `yaml:"arguments" `
	Modifier     string   `yaml:"modifiers"`
	Day          string   `yaml:"day"`
	Month        string   `yaml:"month"`
	StartTime    string   `yaml:"start_time"`
	EndTime      string   `yaml:"end_time"`
	StartDate    string   `yaml:"start_date"`
	EndDate      string   `yaml:"end_date"`
}

type Compose struct {
	Namespace string          `yaml:"namespace"`
	Tasks     map[string]Task `yaml:"tasks"`
}
