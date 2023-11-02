package main

import (
	"github.com/ryanjarv/msh/cmd/common"
	"github.com/ryanjarv/msh/cmd/sfn"
	"github.com/ryanjarv/msh/pkg/app"
	L "github.com/ryanjarv/msh/pkg/logger"
	"log"
	"os"
)

func main() {
	app, err := app.GetPipeline(common.Registry, os.Stdin, os.Stdout)
	if err != nil {
		L.Error.Fatalln("%s: get app: %w", os.Args[0], err)
	}

	l, err := sfn.NewSfn(os.Args)
	if err != nil {
		L.Error.Fatalln("%s: new", l.GetName(), err)
	}

	err = app.Run(l)
	if err != nil {
		log.Fatalf("%s: run: %s", l.GetName(), err)
	}
}
