package main

import (
	"fmt"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/logs"
	"log"
	"os"
	"path/filepath"
)

//go:generate go run ../scripts/gen_modules.go

func main() {
	if err := Run(); err != nil {
		log.Fatalln(err)
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

	if os.Getenv("DEBUG") != "" {
		logs.SetLevel(logs.LevelDebug)
		logs.Debug("debug enabled")
	}

	app, err := app.GetPipeline(Registry, os.Stdin, os.Stdout, os.Args)
	if err != nil {
		return fmt.Errorf("%s: get app: %w", name, err)
	}

	step, err := NewModule(&app, name)
	if err != nil {
		return err
	}

	app.AddStep(step)

	if err = app.Run(); err != nil {
		return fmt.Errorf("run: %s", err)
	}

	return nil
}
