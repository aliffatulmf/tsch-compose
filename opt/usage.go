package opt

import (
	"fmt"
	"os"

	"slices"
)

var layout = `Usage: tsch-compose <command> [options]
Commands:
    %-20s	Run the tasks specified in the configuration file
    %-20s	Show this help message and exit
Options:
    %-20s	Show verbose output

Use 'tsch-compose --help' for more information.
`

type Help struct {
	Args   []string
	action []string
}

func NewHelp() *Help {
	return &Help{
		Args: os.Args[1:],
	}
}

func (h *Help) Action(names []string, fn func() error) {
	for _, name := range names {
		h.action = append(h.action, name)

		if slices.Contains(h.Args, name) {
			err := fn()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func (h *Help) Registered() []string {
	return h.action
}

func (h *Help) CheckInvalid() {
	for _, arg := range h.Args {
		if !slices.Contains(h.action, arg) {
			fmt.Printf("Error: Invalid argument or command: %s\n", arg)
			h.Show()
		}
	}
}

func (h *Help) CheckArgs() {
	argMap := make(map[string]int)

	if len(h.Args) == 0 {
		h.Show()
		return
	}

	for _, arg := range h.Args {
		argMap[arg]++
		if argMap[arg] > 1 {
			fmt.Printf("Error: Duplicate argument or command: %s\n", arg)
			os.Exit(1)
		}
	}
}

func (h *Help) Show() {
	fmt.Printf(layout,
		"up, --up",
		"help, -h, --help",
		"-v, --verbose",
	)
}
