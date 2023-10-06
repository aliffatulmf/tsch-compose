package model

import (
	"fmt"
	"time"

	"github.com/aliffatulmf/tsch-compose/internal/schedule"
)

type Compose struct {
	Namespace string          `yaml:"namespace" validate:"required"`
	Tasks     map[string]Task `yaml:"tasks" validate:"required"`
}

type Time struct {
	Hour   uint `yaml:"hour" validate:"required,lt=23"`
	Minute uint `yaml:"minute" validate:"required,lt=59"`
	Second uint `yaml:"second" validate:"lt=59"`
}

func (t Time) Clock() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
}

type Date struct {
	Month time.Month `yaml:"month" validate:"gte=1,lte=12"`
	Day   uint       `yaml:"day" validate:"gte=1,lte=31"`
	Year  uint       `yaml:"year" validate:"gte=1970,lte=9999"`
}

type Task struct {
	Name         string      `yaml:"name" validate:"required"`
	ScheduleType schedule.ST `yaml:"schedule_type" validate:"required"`
	Run          string      `yaml:"run" validate:"required"`
	Arguments    []string    `yaml:"arguments" `
	Modifier     interface{} `yaml:"modifiers"`
	Day          interface{} `yaml:"day"`
	Month        string      `yaml:"month"`
	StartTime    Time        `yaml:"start_time"`
	EndTime      Time        `yaml:"end_time"`
	StartDate    Date        `yaml:"start_date"`
	EndDate      Date        `yaml:"end_date"`
}
