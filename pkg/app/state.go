package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/types"
	"io"
	"reflect"
)

func ReadState(f io.Reader, registry types.Registry) (State, error) {
	buf := bufio.NewReader(f)

	line, err := buf.ReadBytes('\n')
	if err != nil {
		return State{}, fmt.Errorf("readConf: failed to read line: %w", err)
	}

	return UnmarshalState(registry, line)
}

func UnmarshalState(registry types.Registry, line []byte) (State, error) {
	state := struct {
		Steps []struct {
			Name  string
			Value json.RawMessage
		}
	}{
		Steps: []struct {
			Name  string
			Value json.RawMessage
		}{},
	}

	err := json.Unmarshal(line, &state)
	if err != nil {
		return State{}, fmt.Errorf("readConf: failed to unmarshal: %w", err)
	}

	steps := []types.Step{}

	// Set the correct type for each Step.
	for _, step := range state.Steps {
		cmd, ok := registry[step.Name]
		if !ok {
			return State{}, fmt.Errorf("readConf: unknown step: %s", step.Name)
		}

		typ := reflect.TypeOf(cmd).Elem()
		value := reflect.New(typ).Interface()

		err := json.Unmarshal(step.Value, &value)
		if err != nil {
			return State{}, fmt.Errorf("readConf: failed to unmarshal step: %w", err)
		}

		cdkStep, ok := value.(types.CdkStep)
		if !ok {
			return State{}, fmt.Errorf("readConf: step must implement CdkStep: %T", value)
		}

		steps = append(steps, types.Step{
			Name:  step.Name,
			Value: cdkStep,
		})
	}

	return State{
		Steps: steps,
	}, nil
}

type State struct {
	Steps []types.Step
}

// AddStep is called by each individual Step that wants to add to the cumulative app.
func (s *State) AddStep(c types.CdkStep) {
	L.Debug.Println("State.AddStep: %T", c)
	s.Steps = append(s.Steps, types.Step{
		Name:  c.GetName(),
		Value: c,
	})
}
