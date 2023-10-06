package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aliffatulmf/tsch-compose/internal/model"
	"github.com/aliffatulmf/tsch-compose/internal/parser"

	"github.com/aliffatulmf/tsch-compose/cmd"
	"github.com/aliffatulmf/tsch-compose/config"
)

func handleArgument() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "check":
			_, err := config.ReadCompose()
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
			fmt.Println("tsch-compose.yaml is valid")
		case "help":
			fmt.Print(cmd.Usage)
		default:
			fmt.Println("Unknown command")
		}
	}
}

func main() {
	compose := model.OpenComposeFile()

	parser := parser.NewParser(compose)
	parser.Parse("create")
	cmds := parser.ToCMD()

	for name, cmd := range cmds {
		fmt.Printf("Creating task %s\n", name)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
