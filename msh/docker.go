package msh

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	L "github.com/ryanjarv/msh/logger"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
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

func Exec(path string, args []string) {
	L.Debug.Printf("Running Exec with path '%s' and args '%v'\n", path, args)

	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	params := ExecParams{
		Path: absPath,
		Name: filepath.Base(absPath),
	}

	// Assume context directory where the Dockerfile is.
	if tmpDir, err := ioutil.TempDir(os.TempDir(), "msh"); err != nil {
		panic(err)
	} else {
		params.ContextDir = tmpDir
		defer func() {
			if err := os.RemoveAll(tmpDir); err != nil {
				panic(err)
			}
		}()
	}

	var fmtArgs []string
	for _, arg := range args {
		tmpl, err := template.New("path").Parse(arg)
		if err != nil {
			panic(err)
		}

		w := bytes.NewBuffer([]byte{})
		if err := tmpl.Execute(w, params); err != nil {
			panic(err)
		}

		fmtArgs = append(fmtArgs, w.String())
	}

	L.Debug.Printf("Running: %v", fmtArgs)
	cmd := exec.Command(fmtArgs[0], fmtArgs[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Env = append(cmd.Env, os.Environ()...)

	// We look for the app arg separator and start adding shell style environment
	// variables to make available to the compose file.
	if result := InSliceStr(args, "--"); result != nil {
		for i, arg := range args[*result:] {
			cmd.Env = append(cmd.Env, fmt.Sprintf("ARG_%d=%s", i, arg))
		}
	}

	log.Printf("Running command and waiting for it to finish...")
	if err := cmd.Run(); err != nil {
		cmd.StderrPipe()
		panic(err)
	}
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
		Exec(path, []string{
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
		Exec(path, []string{"docker", "build", "-q", "-t", "{{.Name}}", "-f", "{{.Path}}", "{{.ContextDir}}"})
	} else {
		Exec(path, []string{"docker", "build", "-t", "{{.Name}}", "-f", "{{.Path}}", "{{.ContextDir}}"})
	}
	Exec(path, append([]string{"docker", "run", "-i", "{{.Name}}", "--"}, argv[1:]...))
}
