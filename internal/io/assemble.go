package io

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func NormalString(s string) string {
	s = strings.ToLower(s)
	return regexp.MustCompile(`\s+`).ReplaceAllString(s, "_")
}

func (c *Compose) Parse() map[string][]string {
	var dict = make(map[string][]string)

	for name, task := range c.Tasks {
		dict[name] = []string{"-Command", "SCHTASKS", "/CREATE"}

		if task.Name != "" {
			dict[name] = append(dict[name], "/TN", fmt.Sprintf(`"%s\%s"`, c.Namespace, task.Name))
		}

		if task.ScheduleType != "" {
			dict[name] = append(dict[name], "/SC", task.ScheduleType)
		}

		if task.Run != "" && task.Arguments != nil {
			args := strings.Join(task.Arguments, ";")
			dict[name] = append(dict[name], "/TR", fmt.Sprintf(`"%s %s"`, task.Run, args))
		} else if task.Run != "" && task.Arguments == nil {
			dict[name] = append(dict[name], "/TR", task.Run)
		}

		if task.Modifier != "" {
			dict[name] = append(dict[name], "/MO", task.Modifier)
		}

		if task.Month != "" {
			dict[name] = append(dict[name], "/M", task.Month)
		}

		if task.Day != "" {
			dict[name] = append(dict[name], "/D", task.Day)
		}

		if task.StartTime != "" {
			dict[name] = append(dict[name], "/ST", task.StartTime)
		}

		if task.EndTime != "" {
			dict[name] = append(dict[name], "/ET", task.EndTime)
		}

		if task.StartDate != "" {
			dict[name] = append(dict[name], "/SD", task.StartDate)
		}

		if task.EndDate != "" {
			dict[name] = append(dict[name], "/ED", task.EndDate)
		}

		// Create task even if it's already exist
		dict[name] = append(dict[name], "/F")
	}

	return dict
}

func (c *Compose) Command() map[string]*exec.Cmd {
	var cmd = make(map[string]*exec.Cmd)
	args := c.Parse()

	for name := range c.Tasks {
		cmd[name] = exec.Command("powershell.exe", args[name]...)
	}

	return cmd
}
