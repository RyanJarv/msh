package msh

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

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	params := ExecParams{
		Path: absPath,
		Name: filepath.Base(absPath),
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
		if err != nil {
			return err
		}

		w := bytes.NewBuffer([]byte{})
		if err := tmpl.Execute(w, params); err != nil {
			return err
		}

		fmtArgs = append(fmtArgs, w.String())
	}

	L.Debug.Printf("Running: %v", fmtArgs)
	cmd := exec.Command(fmtArgs[0], fmtArgs[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Env = append(cmd.Env, os.Environ()...)

	// We look for the app arg separator and start adding shell style environment
	// variables to make available to the compose file.
	subCmdArgs := *InSliceStr(args, "--")
	for i, arg := range args[subCmdArgs:] {
		cmd.Env = append(cmd.Env, fmt.Sprintf("ARG_%d=%s", i, arg))
	}

	log.Printf("Running command and waiting for it to finish...")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Dockerfile(argv []string) error {
	path := argv[0]

	if os.Getenv("MSH_CONTEXT_TRIGGER_DEPLOY") == "true" {
		return Exec(path, []string{
			"copilot", "task", "run",
			"--default",
			"--follow",
			"--dockerfile", "{{.Path}}",
			fmt.Sprintf(`--command='%s'`, strings.Join(argv[1:], " ")),
		})
	} else {
		err := Exec(path, []string{ "lib", "build", "-t", "{{.name}}", "-f", "{{.Path}}", "{{.ContextDir}}"})
		if err != nil {
			return err
		}
		return Exec(path, append([]string{"lib", "run", "-i", "{{.name}}", "--"}, argv[1:]...))
	}
}

func Compose(argv []string) error {
	cmdStr := append([]string{"lib-compose", "-f", "{{.Path}}", "run", "app", "--"}, argv[1:]...)
	return Exec(argv[0], cmdStr)
}