package main

import (
	"github.com/ryanjarv/msh/cmd/common"
	sfn_filter "github.com/ryanjarv/msh/cmd/sfn.filter"
	"github.com/ryanjarv/msh/pkg/app"
	L "github.com/ryanjarv/msh/pkg/logger"
	"log"
	"os"
)

func main() {
	app, err := app.GetPipeline(common.Registry, os.Stdin, os.Stdout)
	if err != nil {
		L.Error.Fatalf("filter get pipeline: %s", err)
	}

	l, err := sfn_filter.New(os.Args)
	if err != nil {
		L.Error.Fatalf("filter new: %s", err)
	}

	err = app.Run(l)
	if err != nil {
		log.Fatalf("filter run: %s", err)
	}
}
