package main

import (
	"flag"
	"github.com/ryanjarv/msh/cmd/lambda"
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

	l, err := lambda.NewLambda(flag.Args())
	if err != nil {
		L.Error.Fatalln("failed to create lambda", err)
	}

	err = app.Run(l)
	if err != nil {
		log.Fatalf("run: failed to run app: %s", err)
	}
}
