package main

import (
	"flag"
	"github.com/ryanjarv/msh/cmd/common"
	"github.com/ryanjarv/msh/cmd/sleep"
	"github.com/ryanjarv/msh/pkg/app"
	L "github.com/ryanjarv/msh/pkg/logger"
	"log"
	"os"
)

func main() {
	flag.Parse()
	L.Debug.Println("args:", flag.Args())

	app, err := app.GetPipeline(common.Registry, os.Stdin, os.Stdout)
	if err != nil {
		L.Error.Fatalf("sleep get app: %s", err)
	}

	l, err := sleep.NewSleep(flag.Args())
	if err != nil {
		L.Error.Fatalf("sleep: %s", err)
	}

	if err := app.Run(l); err != nil {
		log.Fatalf("sleep: run: %s", err)
	}
}
