package main

import (
	"flag"
	"github.com/ryanjarv/msh/cmd/api"
	"github.com/ryanjarv/msh/cmd/common"
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

	l, err := api.New()
	if err != nil {
		L.Error.Fatalln("api new: %w", err)
	}

	err = app.Run(l)
	if err != nil {
		log.Fatalf("api run: %w", err)
	}
}
