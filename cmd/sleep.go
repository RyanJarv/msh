package main

import (
	"flag"
	"github.com/ryanjarv/msh/cmd/sleep"
	"github.com/ryanjarv/msh/pkg/app"
	L "github.com/ryanjarv/msh/pkg/logger"
	"log"
	"os"
)

func main() {
	flag.Parse()
	L.Debug.Println("args:", flag.Args())

	app, err := app.GetApp(os.Stdin)
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
