package utils

import (
	"fmt"
	L "github.com/ryanjarv/msh/pkg/logger"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
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
	if in.Stdin != nil {
		cmd.Stdin = in.Stdin
	}
	if in.Stdout != nil {
		cmd.Stdout = in.Stdout
	}
	if in.Stderr != nil {
		cmd.Stderr = in.Stderr
	}

	cmd.Env = append(cmd.Env, os.Environ()...)

	// We look for the app arg separator and start adding shell style environment
	// variables to make available to the compose file.
	if result := InSliceStr(in.Args, "--"); result != nil {
		for i, arg := range in.Args[*result:] {
			cmd.Env = append(cmd.Env, fmt.Sprintf("ARG_%d=%s", i, arg))
		}
	}

	L.Debug.Printf("running: %s %s", in.Cmd, strings.Join(in.Args, " "))
	err := cmd.Run()
	return Wrap(err, "done")
}

// InSliceStr will return the index of str in slice if it exists, otherwise it will return nil
func InSliceStr(slice []string, str string) *int {
	for i, val := range slice {
		if val == str {
			return &i
		}
	}
	return nil
}

// IsLeader returns true if the current process is the current process group leader.
//
// The leader can be determined by checking if the process group id is the same as the process id
// and means that we are the last process in the pipeline.
func IsLeader() bool {
	pgid, err := syscall.Getpgid(os.Getpid())
	if err != nil {
		L.Error.Fatalln("failed to get pgid", err)
	}

	if pid := os.Getpid(); pid == pgid {
		return true
	}
	return false
}
