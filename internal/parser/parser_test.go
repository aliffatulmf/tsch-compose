package parser_test

import (
	"testing"

	"github.com/aliffatulmf/tsch-compose/internal/model"
	"github.com/aliffatulmf/tsch-compose/internal/parser"
)

func TestNewParser(t *testing.T) {
	compose := model.OpenComposeFile()

	parser := parser.NewParser(compose)
	args := parser.Parse("create")
	cmds := parser.ToCMD()

	t.Log(args)
	for name, cmd := range cmds {
		t.Log("executing task", name)
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}
	}
}
