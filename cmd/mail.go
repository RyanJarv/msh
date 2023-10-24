package main

import (
	"flag"
	"github.com/ryanjarv/msh/cmd/mail"
	"github.com/ryanjarv/msh/pkg/app"
	L "github.com/ryanjarv/msh/pkg/logger"
	"log"
	"os"
)

func main() {
	flag.Parse()

	app, err := app.GetApp(os.Stdin)
	if err != nil {
		L.Error.Fatalln("failed to read config", err)
	}

	l, err := mail.NewMail()
	if err != nil {
		L.Error.Fatalln("failed to create event", err)
	}

	err = app.Run(l)
	if err != nil {
		log.Fatalf("run: failed to run app: %s", err)
	}
}
