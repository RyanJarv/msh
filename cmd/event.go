package main

import (
	"flag"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/event"
	"github.com/ryanjarv/msh/pkg/state"
	"log"
	"os"
)

func main() {
	flag.Parse()

	app, err := state.GetApp(os.Stdin)
	if err != nil {
		L.Error.Fatalln("failed to read config", err)
	}

	l, err := event.NewEvent()
	if err != nil {
		L.Error.Fatalln("failed to create event", err)
	}

	err = app.Run(l)
	if err != nil {
		log.Fatalf("run: failed to run app: %s", err)
	}
}
