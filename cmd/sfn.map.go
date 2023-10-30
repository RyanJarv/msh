package main

import (
	"flag"
	"github.com/ryanjarv/msh/cmd/common"
	sfn_map "github.com/ryanjarv/msh/cmd/sfn.map"
	"github.com/ryanjarv/msh/pkg/app"
	L "github.com/ryanjarv/msh/pkg/logger"
	"log"
	"os"
)

func main() {
	flag.Parse()

	app, err := app.GetPipeline(common.Registry, os.Stdin, os.Stdout)
	if err != nil {
		L.Error.Fatalln("map: get app:", err)
	}

	l, err := sfn_map.New(app.Registry)
	if err != nil {
		L.Error.Fatalln("map:", err)
	}

	err = app.Run(l)
	if err != nil {
		log.Fatalln("map: run:", err)
	}
}
