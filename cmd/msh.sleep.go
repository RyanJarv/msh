package main

import (
	"flag"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/sleep"
	"github.com/ryanjarv/msh/pkg/state"
	"log"
	"os"
)

func main() {
	flag.Parse()
	L.Debug.Println("args:", flag.Args())

	app, err := state.GetApp(os.Stdin)
	if err != nil {
		L.Error.Fatalln("failed to read config", err)
	}

	l, err := sleep.NewSleep(flag.Args())
	if err != nil {
		L.Error.Fatalln("failed to create lambda", err)
	}

	if err := app.Run(l); err != nil {
		log.Fatalf("run: failed to run app: %s", err)
	}
}
