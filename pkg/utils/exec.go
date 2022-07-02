package utils

import (
	"fmt"
	L "github.com/ryanjarv/msh/logger"
	"io"
	"os"
	"os/exec"
)

type ExecuteInput struct {
	Cmd    string
	Args   []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func Execute(in *ExecuteInput) error {
	cmd := exec.Command(in.Cmd, in.Args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Env = append(cmd.Env, os.Environ()...)

	// We look for the app arg separator and start adding shell style environment
	// variables to make available to the compose file.
	if result := InSliceStr(in.Args, "--"); result != nil {
		for i, arg := range in.Args[*result:] {
			cmd.Env = append(cmd.Env, fmt.Sprintf("ARG_%d=%s", i, arg))
		}
	}

	L.Debug.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	return Wrap(err, "running cmd")
}
