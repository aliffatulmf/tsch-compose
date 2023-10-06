package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var ScheduleTypeArray = [9]string{"MINUTE", "HOURLY", "DAILY", "WEEKLY", "MONTHLY", "ONCE", "ONSTART", "ONLOGON", "ONIDLE"}

// Specify custom validator for Config with sturct validation enabled
func Validator(c *Config, tags ...func(v *validator.Validate) error) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if len(tags) > 0 {
		for _, f := range tags {
			if err := f(validate); err != nil {
				return err
			}
		}
	}

	if err := validate.Struct(c); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%s: %s", err.Namespace(), err.Tag())
		}
	}

	return nil
}
