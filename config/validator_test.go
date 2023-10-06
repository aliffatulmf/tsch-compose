package config_test

import (
	"net"
	"slices"
	"testing"

	"github.com/aliffatulmf/tsch-compose/config"
	"github.com/go-playground/validator/v10"
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

func TestConfigValidator_IncorrectScheduleType(t *testing.T) {
	validationError := config.Validator(c, func(v *validator.Validate) error {
		return v.RegisterValidation("scheduleType", func(fl validator.FieldLevel) bool {
			return slices.Contains(config.ScheduleTypeArray[:], fl.Field().String())
		})
	})

	if validationError != nil {
		t.Log(validationError)
	} else {
		t.Fatal("expected error, got nil")
	}
}

func TestIP(t *testing.T) {
	ip := net.IPv4(192, 168, 1, 1)

	t.Log(ip.To4())
}
