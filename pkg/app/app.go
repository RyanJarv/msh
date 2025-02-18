package app

import (
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/types"
	"os"
	"strings"
)

func GetPipeline(reg Registry, stdin *os.File, stdout *os.File, args []string) (App, error) {
	L.Debug.Printf("GetPipeline: %s:", args)

	return App{
		Stdin:    stdin,
		Stdout:   stdout,
		Registry: reg,
	}, nil
}

type App struct {
	Name     string
	Steps    []types.Step
	Registry Registry `json:"-"`
	Stdin    *os.File `json:"-"`
	Stdout   *os.File `json:"-"`
}

func (s *App) StackName() string {
	name := s.Name

	name = strings.Replace(name, " ", "-", -1)
	name = strings.Replace(name, "_", "-", -1)
	name = strings.ToLower(name)

	return name
}

// AddStep is called by each individual Step that wants to add to the cumulative app.
func (s *App) AddStep(c types.CdkStep) {
	L.Debug.Println("State.AddStep: %T", c)
	s.Steps = append(s.Steps, types.Step{
		Name:  c.GetName(),
		Value: c,
	})
}

type Registry map[string]func(*App, []string) (types.CdkStep, error)

func (r Registry) Get(name string) func(*App, []string) (types.CdkStep, error) {
	if v, ok := r[name]; ok {
		return v
	}
	return nil
}
