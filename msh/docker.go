package msh

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

const MSH_CONTEXT = "MSH_CONTEXT"

type RunEnv string
const EcsRunEnv RunEnv = "ecs"

type MshContext struct {
	RunEnv RunEnv
}

type ExecParams struct {
	Path       string
	ContextDir string
	Name       string
}


func Dockerfile(argv []string) {
	path := argv[0]

	if c := GetMshContext(); c != nil {
		DockerfileRemote(c.RunEnv, path, argv)
	} else {
		DockerfileLocal(argv, path)
	}
}

func DockerfileRemote(env RunEnv, path string, args []string) {
	switch env {
	case EcsRunEnv:
		DockerTmplExecNoContext(path, []string{
			"copilot", "task", "run",
			"--default",
			"--follow",
			"--dockerfile", "{{.Path}}",
			fmt.Sprintf(`--command='%s'`, strings.Join(args[1:], " ")),
		})
	default:
		panic(errors.Errorf("unknown run environment: '%s'", env))
	}
}

func DockerfileLocal(argv []string, path string) {
	if os.Getenv("DEBUG") == "" {
		DockerTmplExecNoContext(path, []string{"docker", "build", "-q", "-t", "{{.Name}}", "-f", "{{.Path}}", "{{.ContextDir}}"})
	} else {
		DockerTmplExecNoContext(path, []string{"docker", "build", "-t", "{{.Name}}", "-f", "{{.Path}}", "{{.ContextDir}}"})
	}
	DockerTmplExecNoContext(path, append([]string{"docker", "run", "-i", "{{.Name}}", "--"}, argv[1:]...))
}
