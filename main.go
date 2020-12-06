package main

import (
	"flag"
	"github.com/pkg/errors"
	"github.com/ryanjarv/msh/lib"
	L "github.com/ryanjarv/msh/logger"
	"log"
	"os"
)

//go:generate go run script/templates.go


type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	if err := run(); err != nil {
		if err, ok := err.(stackTracer); ok {
			for _, f := range err.StackTrace() {
				L.Error.Printf("%+s:%d\n", f, f)
			}
		}
		log.Fatalln(err)
	}
}

var args = flag.CommandLine

func run() error {
	if err := args.Parse(os.Args[1:]); err != nil {
		return errors.WithStack(err)
	}

	L.Debug.Println("args: %v", args.Args())

	var err error
	switch cmd := args.Args()[0]; cmd {
	case "lambda":
		err = lib.Lambda(args.Args()[1:])
	case "compose":
		err = lib.Compose(args.Args()[1:])
	case "dockerfile":
		err = lib.Dockerfile(args.Args()[1:])
	case "ecs":
		err = lib.Ecs(args.Args()[1:])
	case "swf":
		err = lib.Swf(args.Args()[1:])
	case "cfn":
		err = lib.Cfn(args.Args()[1:])
	case "init":
		err = lib.Init(args.Args()[1:])
	case "":
		args.Usage()
		err = errors.New("Must specify subcommand")
	default:
		return errors.Errorf("unknown command: '%v'", cmd)
	}
	return err
}

