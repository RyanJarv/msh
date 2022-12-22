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
	lambda := lambda.NewLambdaCmd(
		command.NewCommand(flag.Args()),
	)
	err := lambda.Deploy()
	if err != nil {
		L.Error.Fatalln("failed to deploy lambda:", err)
	}

	proc := process.NewMsh(lambda)

	err = proc.Run()
	if err != nil {
		L.Error.Fatalln(err)
	}
}
