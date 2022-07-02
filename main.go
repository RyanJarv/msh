package main

import (
	"flag"
	"fmt"
	"github.com/ryanjarv/msh/pkg"
	L "github.com/ryanjarv/msh/pkg/logger"
	"os"
)

func main() {
	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		fmt.Printf("Usage: %s <path>\n", os.Args[0])
		os.Exit(1)
	}

	progArgs := flag.Args()[1:]

	L.Debug.Printf("Arguments: cmd: %s, path: %s, progArgs: %+v", path, progArgs)

	err := pkg.Dockerfile(path, progArgs...)

	if err != nil {
		L.Error.Fatalf("[ERROR] %s\n", err)
	}
}
