package main

import (
	"flag"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/command"
	"github.com/ryanjarv/msh/pkg/providers/process"
)

func main() {
	flag.Parse()
	cmd := command.NewCommand(flag.Args())
	proc := process.NewProcess(cmd)

	err := proc.Run()
	if err != nil {
		L.Error.Fatalln(err)
	}
}
