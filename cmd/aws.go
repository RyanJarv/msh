package main

import (
	"github.com/ryanjarv/msh/cmd/aws"
	"github.com/ryanjarv/msh/cmd/common"
	"github.com/ryanjarv/msh/pkg/app"
	L "github.com/ryanjarv/msh/pkg/logger"
	"log"
	"os"
)

func main() {
	app, err := app.GetPipeline(common.Registry, os.Stdin, os.Stdout)
	if err != nil {
		L.Error.Fatalln("aws get app:", err)
	}

	l, err := aws.New(os.Args)
	if err != nil {
		L.Error.Fatalln("aws new:", err)
	}

	err = app.Run(l)
	if err != nil {
		log.Fatalln("aws run:", err)
	}
}
