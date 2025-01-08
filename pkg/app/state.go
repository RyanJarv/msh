package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/types"
	"io"
	"os"
	"reflect"
	"strings"
)

type State struct {
	Name  string
	Steps []types.Step
}

func (s *State) StackName() string {
	name := s.Name

	name = strings.Replace(name, " ", "-", -1)
	name = strings.Replace(name, "_", "-", -1)
	name = strings.ToLower(name)

	return name
}

// AddStep is called by each individual Step that wants to add to the cumulative app.
func (s *State) AddStep(c types.CdkStep) {
	L.Debug.Println("State.AddStep: %T", c)
	s.Steps = append(s.Steps, types.Step{
		Name:  c.GetName(),
		Value: c,
	})
}

func (s *State) WriteState(f *os.File) error {
	d, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("writeState: failed to marshal app: %w", err)
	}

	L.Debug.Printf("State.WriteState: writing to %s: %s", f.Name(), string(d))
	_, err = io.WriteString(f, string(d)+"\n")
	if err != nil {
		return fmt.Errorf("writeState: failed to write app: %w", err)
	}

	return nil
}

func ReadState(f io.Reader, registry types.Registry) (*State, error) {
	buf := bufio.NewReader(f)

	line, err := buf.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("readConf: failed to read line: %w", err)
	}

	return UnmarshalState(registry, line)
}

func UnmarshalState(registry types.Registry, line []byte) (*State, error) {
	state := struct {
		Name  string
		Steps []struct {
			Name  string
			Value json.RawMessage
		}
	}{
		Name: "msh-default",
		Steps: []struct {
			Name  string
			Value json.RawMessage
		}{},
	}

	err := json.Unmarshal(line, &state)
	if err != nil {
		return nil, fmt.Errorf("readConf: failed to unmarshal: %w", err)
	}

	steps := []types.Step{}

	// Set the correct type for each Step.
	for _, step := range state.Steps {
		cmd, ok := registry[step.Name]
		if !ok {
			return nil, fmt.Errorf("readConf: unknown this: %s", step.Name)
		}

		typ := reflect.TypeOf(cmd).Elem()
		value := reflect.New(typ).Interface()

		err := json.Unmarshal(step.Value, &value)
		if err != nil {
			return nil, fmt.Errorf("readConf: failed to unmarshal this: %w", err)
		}

		cdkStep, ok := value.(types.CdkStep)
		if !ok {
			return nil, fmt.Errorf("readConf: this must implement CdkStep: %T", value)
		}

		steps = append(steps, types.Step{
			Name:  step.Name,
			Value: cdkStep,
		})
	}

	return &State{
		Name:  state.Name,
		Steps: steps,
	}, nil
}
