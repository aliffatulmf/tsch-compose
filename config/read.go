package config

import (
	"errors"
	"fmt"
	"net"
	"os"
	"slices"

	"github.com/go-playground/validator/v10"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Namespace string          `yaml:"namespace" validate:"required"`
	Tasks     map[string]Task `yaml:"tasks" validate:"required,dive"`
}

type Task struct {
	Name         string   `yaml:"name" validate:"required"`
	ScheduleType string   `yaml:"schedule_type" validate:"required,scheduleType"`
	Run          string   `yaml:"run"`
	Arguments    []string `yaml:"arguments"`
	Remote       net.IP   `yaml:"remote"`
	Modifiers    []string `yaml:"modifiers"`
}

const FileName = "tsch-compose.yaml"

func read(cfg *Config) error {
	cfgByte, err := os.ReadFile(FileName)
	if err != nil {
		return errors.New("no tsch-compose.yaml found")
	}

	if err := yaml.Unmarshal(cfgByte, cfg); err != nil {
		return err
	}

	if err := Validator(cfg, func(v *validator.Validate) error {
		return v.RegisterValidation("scheduleType", func(fl validator.FieldLevel) bool {
			return slices.Contains(ScheduleTypeArray[:], fl.Field().String())
		})
	}, func(v *validator.Validate) error {
		err := v.Struct(cfg)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				return fmt.Errorf("got error on field %s: %s", err.Field(), err.Tag())
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func ReadCompose() (*Config, error) {
	cfg := &Config{}
	if err := read(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
