package lib

import (
	"bytes"
	"fmt"
	L "github.com/ryanjarv/msh/logger"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

type ExecParams struct {
	Path       string
	ContextDir string
	Name       string
}

func Exec(path string, args []string) error {
	L.Debug.Printf("Running Exec with path '%s' and args '%v'\n", path, args)

	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	params := ExecParams{
		Path: path,
		Name: filepath.Base(path),
	}

	if tmpDir, err := ioutil.TempDir(os.TempDir(), "msh"); err != nil {
		return err
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
		if err != nil { return err }

		w := bytes.NewBuffer([]byte{})
		err = tmpl.Execute(w, params)
		if err != nil {
			return err
		}

		fmtArgs = append(fmtArgs, w.String())
	}

	L.Debug.Printf("Running: %v", fmtArgs)
	cmd := exec.Command(fmtArgs[0], fmtArgs[1:]...)

	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	log.Printf("Running command and waiting for it to finish...")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Dockerfile(argv []string) error {
	path := argv[0]

	if os.Getenv("MSH_CONTEXT") == "ecs" {
		return Exec(path, []string{
			"copilot", "task", "run",
			"--default",
			"--follow",
			"--dockerfile", "{{.Path}}",
			fmt.Sprintf(`--command='%s'`, strings.Join(argv[1:], " ")),
		})
	} else {
		err := Exec(path, []string{ "docker", "build", "-t", "{{.Name}}", "-f", "{{.Path}}", "{{.ContextDir}}"})
		if err != nil {
			return err
		}
		return Exec(path, append([]string{"docker", "run", "-i", "{{.Name}}"}, argv[1:]...))
	}
}

func Compose(argv []string) error {
	path := argv[0]
	return Exec(path, append([]string{"docker-compose", "-f", "{{.Path}}", "run", "app"}, argv[1:]...))
}

// Ecs sets MSH_CONTEXT=ecs and executes the rest of the args. This typically will be a path to a msh docker config so
// this variable get's checked in Dockerfile.
func Ecs(argv []string) error {
	os.Setenv("MSH_CONTEXT", "ecs")

	cmd := exec.Command(argv[0], argv[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	log.Printf("Running command and waiting for it to finish...")
	return cmd.Run()
}
