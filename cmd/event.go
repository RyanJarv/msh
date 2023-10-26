package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ryanjarv/msh/cmd/common"
	"github.com/ryanjarv/msh/cmd/event"
	"github.com/ryanjarv/msh/pkg/app"
	L "github.com/ryanjarv/msh/pkg/logger"
	acm "github.com/ryanjarv/msh/schema/aws/acm/acmcertificateavailable"
	"github.com/samber/lo"
	"log"
	"os"
)

func main() {
	flag.Parse()

	v := acm.ACMCertificateAvailable{}
	fmt.Println(
		string(lo.Must(json.Marshal(v))),
	)

	app, err := app.GetPipeline(common.Registry, os.Stdin, os.Stdout)
	if err != nil {
		L.Error.Fatalln("%s: get app: %w", os.Args[0], err)
	}

	l, err := event.NewEvent(flag.Args())
	if err != nil {
		L.Error.Fatalln("%s: new", l.GetName(), err)
	}

	err = app.Run(l)
	if err != nil {
		log.Fatalf("%s: run: %s", l.GetName(), err)
	}
}
