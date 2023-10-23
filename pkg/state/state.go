package state

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/utils"
	"os"
)

func GetApp(file *os.File) (App, error) {
	app := App{
		State: &State{
			[]step{},
		},
		registry: SetupRegistry(),
	}
	if utils.IsTTY(file) {
		return app, nil
	}

	buf := bufio.NewReader(file)

	line, err := buf.ReadBytes('\n')
	if err != nil {
		return app, fmt.Errorf("readConf: failed to read line: %w", err)
	}

	app.State, err = app.ReadState(line)
	if err != nil {
		return app, fmt.Errorf("getApp: failed to read state: %w", err)
	}

	return app, nil
}

func (a *App) ReadState(line []byte) (*State, error) {
	state := &State{
		Steps: []step{},
	}

	err := json.Unmarshal(line, state)
	if err != nil {
		return state, fmt.Errorf("readConf: failed to unmarshal: %w", err)
	}

	// Set the correct type for each step.
	for i, step := range state.Steps {
		state.Steps[i].Value = a.registry[step.Name]
	}

	// Unmarshall into the new types.
	err = json.Unmarshal(line, state)
	if err != nil {
		return state, fmt.Errorf("readConf: failed to unmarshal: %w", err)
	}

	return state, nil
}

type App struct {
	State    *State
	registry Registry
}

type step struct {
	Name  string
	Value interface{}
}

type State struct {
	Steps []step
}

// AddStep is called by each individual step that wants to add to the cumulative state.
func (s *State) AddStep(c types.IStep) {
	s.Steps = append(s.Steps, step{
		Name:  c.Name(),
		Value: c,
	})
}
