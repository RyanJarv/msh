package main

import (
	"bytes"
	"flag"
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

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	if err := run(); err != nil {
		if err, ok := err.(stackTracer); ok {
			for _, f := range err.StackTrace() {
				L.Error.Printf("%+s:%d\n", f, f)
			}
		}
		log.Fatalln(err)
	}
}

var args = flag.CommandLine

func run() error {
	if err := args.Parse(os.Args[1:]); err != nil {
		return errors.WithStack(err)
	}

	L.Debug.Println("args: %v", args.Args())

	switch cmd := args.Args()[0]; cmd {
	case "dockerfile":
		Dockerfile(args.Args()[1:])
	case "ecs":
		Ecs(args.Args()[1:])
	default:
		return errors.Errorf("unknown command: '%v'", cmd)
	}
	return nil
}

type ExecParams struct {
	Path       string
	ContextDir string
	Name       string
}
func Exec(path string, args []string) error {
	L.Debug.Printf("Running Exec with path '%s' and args '%v'\n", path, args)

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

// Ecs sets MSH_CONTEXT=ecs and executes the rest of the args. This typically will be a path to a msh docker config so
// this variable get's checked in Dockerfile.
func Ecs(argv []string) error {
	os.Setenv("MSH_CONTEXT", "ecs")

	cmd := exec.Command(argv[0], argv[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	log.Printf("Running command and waiting for it to finish...")
	return cmd.Run()
}
