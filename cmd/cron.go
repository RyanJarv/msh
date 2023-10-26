package main

import (
	"flag"
	"github.com/ryanjarv/msh/cmd/common"
	"github.com/ryanjarv/msh/cmd/cron"
	"github.com/ryanjarv/msh/pkg/app"
	L "github.com/ryanjarv/msh/pkg/logger"
	"os"
)

func main() {
	flag.Parse()

	app, err := app.GetPipeline(common.Registry, os.Stdin, os.Stdout)
	if err != nil {
		L.Error.Fatalf("cron: get app: %w", err)
	}

	l, err := event.NewCron()
	if err != nil {
		L.Error.Fatalf("cron: new: %w", err)
	}

	err = app.Run(l)
	if err != nil {
		L.Error.Fatalf("cron: run: %w", err)
	}
}
