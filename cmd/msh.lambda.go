package main

import (
	"flag"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/lambda"
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

	l, err := lambda.NewLambda(flag.Args())
	if err != nil {
		L.Error.Fatalln("failed to create lambda", err)
	}

	err = app.Run(l)
	if err != nil {
		log.Fatalf("run: failed to run app: %w", err)
	}
}
