package app

import (
	"flag"
	"fmt"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/utils"
	"os"
)

func GetPipeline(reg types.Registry, stdin *os.File, stdout *os.File, args []string) (App, error) {
	var state State

	// We need to read the state before we can do anything else.
	mshStdin := os.Getenv("MSH_STDIN")
	if !utils.IsTTY(stdin) || mshStdin != "" {
		var err error

		// Hack to allow us to read from a test file in a debugger.
		if mshStdin != "" {
			stdin, err = os.Open(mshStdin)
			if err != nil {
				return App{}, fmt.Errorf("run: failed to open stdin: %w", err)
			}
		}

		state, err = ReadState(stdin, reg)
		if err != nil {
			return App{}, fmt.Errorf("run: failed to read state: %w", err)
		}
	} else {
		state = State{
			Steps: []types.Step{},
		}
	}

	return App{
		FlagSet:  utils.ParseArgs(args),
		State:    state,
		Stdin:    stdin,
		Stdout:   stdout,
		Registry: reg,
		OsArgs:   args,
	}, nil
}

type App struct {
	*flag.FlagSet `json:"-"`
	State         State
	Registry      types.Registry `json:"-"`
	OsArgs        []string       `json:"-"`
	Stdin         *os.File       `json:"-"`
	Stdout        *os.File       `json:"-"`
}
