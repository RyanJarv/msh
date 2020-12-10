package main

import (
	"flag"
	"github.com/pkg/errors"
	L "github.com/ryanjarv/msh/logger"
	"github.com/ryanjarv/msh/msh"
	"os"
)

var args = flag.CommandLine

func main() {
	if err := args.Parse(os.Args[1:]); err != nil {
		panic(errors.Wrap(err, "missing argument"))
	}

	L.Debug.Println("args: %v", args.Args())

	switch cmd := args.Args()[0]; cmd {
	case "compose":
		msh.Compose(args.Args()[1:])
	case "dockerfile":
		msh.Dockerfile(args.Args()[1:])
	case "ecs":
		msh.Remote(args.Args()[0:])
	case "":
		args.Usage()
		errors.New("Must specify subcommand")
	default:
		panic(errors.Errorf("unknown command: '%v'", cmd))
	}
}
