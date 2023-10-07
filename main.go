package main

import (
	"fmt"

	tschio "github.com/aliffatulmf/tsch-compose/internal/io"
	"github.com/aliffatulmf/tsch-compose/opt"
)

var help *opt.Help

func init() {
	help = opt.NewHelp()
	help.CheckArgs()
}

func main() {
	executeCommand()
}

func executeCommand() {
	var verbose bool

	help.Action([]string{"-v", "--verbose"}, func() error {
		verbose = true
		return nil
	})
	help.Action([]string{"up"}, func() error {
		RunConfiguredTasks(verbose)
		return nil
	})
	help.Action([]string{"help", "-h", "--help"}, func() error {
		help.Show()
		return nil
	})

	help.CheckInvalid()
}

func RunConfiguredTasks(verbose bool) {
	compose := tschio.ReadFile()

	for _, cmd := range compose.Command() {
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(output))
			return
		}

		if verbose {
			fmt.Printf("%s\r", string(output))
		}
	}
}
