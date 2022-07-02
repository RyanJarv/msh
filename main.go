package main

import (
	"flag"
	"fmt"
	L "github.com/ryanjarv/msh/logger"
	"github.com/ryanjarv/msh/pkg"
	"os"
)

var args = flag.CommandLine

//go:generate go run ./scripts/templates.go

func main() {
	flag.Parse()

	cmd := flag.Arg(0)
	path := flag.Arg(1)
	if cmd == "" || path == "" {
		fmt.Printf("Usage: %s <cmd> <path>\n", os.Args[0])
		os.Exit(1)
	}

	progArgs := flag.Args()[2:]

	L.Debug.Printf("Arguments: cmd: %s, path: %s, progArgs: %+v", cmd, path, progArgs)

	err := pkg.Dockerfile(path, progArgs...)

	if err != nil {
		L.Error.Fatalf("[ERROR] %s\n", err)
	}
}
