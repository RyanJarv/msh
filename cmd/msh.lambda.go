package main

import (
	"github.com/ryanjarv/msh/pkg/providers/pipes"
	"github.com/ryanjarv/msh/pkg/providers/process"
)

func main() {
	awsPipe := pipes.NewPipe()
	proc := process.NewProcess(awsPipe)

	err := proc.Run()
	if err != nil {
		panic(err)
	}
}
