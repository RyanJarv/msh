package main

import (
	"flag"
	"github.com/ryanjarv/msh/cmd/common"
	sfn_filter "github.com/ryanjarv/msh/cmd/sfn.filter"
	"github.com/ryanjarv/msh/pkg/app"
	L "github.com/ryanjarv/msh/pkg/logger"
	"log"
	"os"
)

func main() {
	flag.Parse()

	app, err := app.GetPipeline(common.Registry, os.Stdin, os.Stdout)
	if err != nil {
		L.Error.Fatalln("%s: get app: %w", os.Args[0], err)
	}

	l, err := sfn_filter.New(app.Registry)
	if err != nil {
		L.Error.Fatalln("%s: new", l.GetName(), err)
	}

	err = app.Run(l)
	if err != nil {
		log.Fatalf("%s: run: %s", l.GetName(), err)
	}
}
