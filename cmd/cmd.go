package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

type Command struct {
	tasks map[string][]string
}

func NewCommand(tasks map[string][]string) *Command {
	return &Command{
		tasks: tasks,
	}
}

func (c *Command) Run() error {
	for name, args := range c.tasks {
		fmt.Println("creating task", name)

		cmd := exec.Command("powershell.exe", "-Command", "schtasks", strings.Join(args, " "))
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
