package app

import (
	"fmt"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/utils"
	"os"
)

func GetApp(reg types.Registry, stdin *os.File, stdout *os.File) (App, error) {
	var state State

	// We need to read the state before we can do anything else.
	if !utils.IsTTY(stdin) {
		var err error
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
		State:    state,
		Stdin:    stdin,
		Stdout:   stdout,
		Registry: reg,
	}, nil
}

type App struct {
	State    State
	Registry types.Registry `json:"-"`
	Stdin    *os.File       `json:"-"`
	Stdout   *os.File       `json:"-"`
}
