package main

import (
	"flag"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/command"
	"github.com/ryanjarv/msh/pkg/providers/eventbridge"
	"github.com/ryanjarv/msh/pkg/providers/lambda"
	"github.com/ryanjarv/msh/pkg/providers/process"
	"github.com/ryanjarv/msh/pkg/utils"
	"os"
)

func main() {
	flag.Parse()

	cmd := command.NewCommand(flag.Args())

	function := lambda.NewLambdaCmd(cmd)

	var proc *process.Process
	if utils.IsTTY(os.Stdin) {
		L.Debug.Println("running in non-streaming mode")
		proc = process.NewProcess(function)
	} else {
		L.Debug.Println("running in streaming mode")
		pipe := eventbridge.NewPipe(function)
		proc = process.NewProcess(&pipe)
	}

	err := proc.Deploy()
	if err != nil {
		L.Error.Fatalln("eventbridge pipes deploy:", err)
	}

	err = proc.Run()
	if err != nil {
		panic(err)
	}
}
