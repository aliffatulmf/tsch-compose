package main

import (
	"fmt"
	"os"
	"strings"

	"slices"

	tschio "github.com/aliffatulmf/tsch-compose/internal/io"
)

var Usage = `Usage: tsch-compose <command> [options]
Commands:
    %-20s	Run the tasks specified in the configuration file
    %-20s	Show this help message and exit
Options:
    %-20s	Show verbose output

Use 'tsch-compose --help' for more information.
`

func printUsage() {
	fmt.Printf(Usage, "up, --up", "help, -h, --help", "-v, --verbose")
}

func executeCommand(command string, verbose bool) {
	switch strings.ToLower(command) {
	case "up", "--up":
		RunConfiguredTasks(verbose)
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Use 'tsch-compose help' for a list of available commands.")
	}
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

func checkDuplicateArgs(args []string) {
	argMap := make(map[string]int)
	for _, arg := range args {
		argMap[arg]++
		if argMap[arg] > 1 {
			fmt.Printf("Error: Duplicate argument or command: %s\n", arg)
			os.Exit(1)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		printUsage()
		return
	}

	checkDuplicateArgs(os.Args[1:])
	verbose := slices.Contains(os.Args, "-v") ||
		slices.Contains(os.Args, "--verbose")

	executeCommand(os.Args[1], verbose)
}
