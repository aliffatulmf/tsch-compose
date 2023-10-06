package cmd

import (
	"github.com/aliffatulmf/tsch-compose/config"
)

var c = &config.Config{
	Namespace: "namespace",
	Tasks: map[string]config.Task{
		"task1": {
			Name:         "task1",
			ScheduleType: "ONSTART",
			Run:          "cmd",
			Arguments:    []string{"/c", "echo", "hello"},
		},
		"task2": {
			Name:         "task2",
			ScheduleType: "INCORRECT",
			Run:          "cmd",
			Arguments:    []string{"/c", "echo", "world"},
		},
	},
}
