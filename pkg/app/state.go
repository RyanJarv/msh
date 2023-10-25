package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ryanjarv/msh/pkg/types"
	"io"
)

func ReadState(f io.Reader, registry types.Registry) (*State, error) {
	buf := bufio.NewReader(f)

	line, err := buf.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("readConf: failed to read line: %w", err)
	}

	return UnmarshalState(registry, line)
}

func UnmarshalState(registry types.Registry, line []byte) (*State, error) {
	state := &State{
		Steps: []step{},
	}

	err := json.Unmarshal(line, state)
	if err != nil {
		return nil, fmt.Errorf("readConf: failed to unmarshal: %w", err)
	}

	// Set the correct type for each step.
	for i, step := range state.Steps {
		state.Steps[i].Value = registry[step.Name]
	}

	// Unmarshall into the new types.
	err = json.Unmarshal(line, state)
	if err != nil {
		return state, fmt.Errorf("readConf: failed to unmarshal: %w", err)
	}

	return state, nil
}

type step struct {
	Name  string
	Value interface{}
}

type State struct {
	Steps []step
}

// AddStep is called by each individual step that wants to add to the cumulative app.
func (s *State) AddStep(c types.IStep) {
	s.Steps = append(s.Steps, step{
		Name:  c.Name(),
		Value: c,
	})
}
