package main

import (
	"flag"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/command"
	"github.com/ryanjarv/msh/pkg/providers/eventbridge"
	"github.com/ryanjarv/msh/pkg/providers/lambda"
	"github.com/ryanjarv/msh/pkg/providers/process"
)

func main() {
	flag.Parse()
	function := lambda.NewLambdaCmd(
		command.NewCommand(flag.Args()),
	)
	pipe := eventbridge.NewPipe(function)

	proc := process.NewProcess(&pipe)

	err := function.Deploy()
	if err != nil {
		L.Error.Fatalln("failed to deploy lambda:", err)
	}

	err = pipe.Deploy()
	if err != nil {
		L.Error.Fatalln("failed to deploy lambda:", err)
	}

	err = proc.Run()
	if err != nil {
		panic(err)
	}
}
