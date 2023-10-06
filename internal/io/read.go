package io

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFile() *Compose {
	file, err := os.ReadFile("tsch-compose.yaml")
	if err != nil {
		log.Fatalf("cannot open tsch-compose.yaml: %v", err)
	}

	compose := &Compose{}
	if err = yaml.Unmarshal(file, compose); err != nil {
		log.Fatalf("cannot decode tsch-compose.yaml: %v", err)
	}

	return compose
}
