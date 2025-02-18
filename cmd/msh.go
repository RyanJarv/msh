package main

import (
	"flag"
	"fmt"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/utils"
	"log"
	"os"
)

//go:generate go run ../scripts/gen_modules.go

func main() {
	if err := Run(); err != nil {
		log.Fatalln(err)
	}
}

var (
	help = flag.Bool("help", false, "print help")
)

func Run() error {
	flag.Parse()

	for _, path := range flag.Args() {
		apps, err := ParseScript(path)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}

		for _, s := range apps {
			if err = s.Run(); err != nil {
				return fmt.Errorf("run: %s", err)
			}
		}
	}

	return nil
}

func ParseScript(path string) ([]*app.App, error) {
	var pipelines []*app.App

	script, err := utils.ParseFile(path)
	if err != nil {
		return nil, err
	}

	for _, pipe := range script.Pipelines {
		if pipe.Key != "msh" {
			return nil, fmt.Errorf("unknown key: %s", pipe.Key)
		}

		fmt.Println(pipe.ParamString())
		app, err := ParsePipeline(pipe.Block, pipe.ParamString())
		if err != nil {
			return nil, fmt.Errorf("%s: %w", pipe.Name(), err)
		}
		pipelines = append(pipelines, app)
	}
	return pipelines, nil
}

func ParsePipeline(pipe []*utils.Step, name string) (*app.App, error) {
	app := &app.App{
		Name:     name,
		Stdin:    os.Stdin,
		Stdout:   os.Stdout,
		Registry: Registry,
	}

	for _, s := range pipe {
		if f := GetStepFunc(s.Key); f != nil {
			step, err := f(app, s.GetArgs())
			if err != nil {
				return nil, fmt.Errorf("step %s: %w", s.Key, err)
			}
			app.AddStep(step)
		}
	}
	return app, nil
}
