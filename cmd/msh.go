package main

import (
	"fmt"
	"github.com/ryanjarv/msh/pkg/app"
	"log"
	"os"
	"path/filepath"
)

//go:generate go run ../scripts/gen_modules.go

func main() {
	err := Run()
	if err != nil {
		log.Fatalf("Run: %s", err)
	}
}

func Run() error {
	var name string
	if v := os.Getenv("MSH_CMD"); v != "" {
		name = v
	} else {
		path, err := os.Executable()
		if err != nil {
			return fmt.Errorf("os.Executable: %w", err)
		}

		name = filepath.Base(path)
	}

	app, err := app.GetPipeline(Registry, os.Stdin, os.Stdout, os.Args)
	if err != nil {
		return fmt.Errorf("%s: get app: %w", name, err)
	}

	m, err := NewModule(app, name)
	if err != nil {
		return err
	}

	err = app.Run(m)
	if err != nil {
		return fmt.Errorf("run: %s", err)
	}

	return nil
}
