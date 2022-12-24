package main

import (
	"flag"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/command"
	"github.com/ryanjarv/msh/pkg/providers/lambda"
	"github.com/ryanjarv/msh/pkg/providers/process"
)

func main() {
	flag.Parse()
	L.Debug.Println("args:", flag.Args())

	function := lambda.NewLambdaCmd(
		command.NewCommand(flag.Args()),
	)
	err := function.Deploy()
	if err != nil {
		L.Error.Fatalln("failed to deploy function:", err)
	}

	proc := process.NewProcess(function)

	err = proc.Run()
	if err != nil {
		L.Error.Fatalln(err)
	}
}
