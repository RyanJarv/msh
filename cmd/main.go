package main

import (
	"bytes"
	"flag"
	"github.com/pkg/errors"
	L "github.com/ryanjarv/msh/logger"
	"io/ioutil"
	"log"
	"mvdan.cc/sh/syntax"
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

	path := args.Arg(args.NArg()-1)
	L.Debug.Println("Config path: ", path)

	args := args.Args()[:args.NArg()-1]
	L.Debug.Printf("Passed Args: %v\n", args)

	if len(args) == 0 {
		flag.Usage()
		return errors.New("no specified command")
	}

	switch cmd := args[0]; cmd {
	case "dockerfile":
		Dockerfile(path, args[1:])
	case "exec":
		Exec(path, args[1:])
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
func Parse(args []string) ([]syntax.Command, error) {
	r := strings.NewReader(strings.Join(args, " "))
	f, err := syntax.NewParser().Parse(r, "")
	if err != nil {
		return nil, err
	}

	var cmds []syntax.Command
	for i, stmt := range f.Stmts {
		L.Debug.Printf("Cmd %d: %-20T - ", i, stmt.Cmd)
		cmds = append(cmds, stmt.Cmd)
	}
	return cmds, nil
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



	cmds, err := Parse(args)
	if err != nil {
		return err
	}

	printer := syntax.NewPrinter()
	for _, cmd := range cmds {
		b := new(bytes.Buffer)
		printer.Print(b, cmd)
		tmpl, err := template.New("path").Parse(b.String())
		if err != nil { return err }

		w := bytes.NewBuffer([]byte{})
		err = tmpl.Execute(w, params)
		if err != nil {
			return err
		}

		cmdArgv := strings.Split(w.String(), " ")
		L.Debug.Printf("Running: %s %s", cmdArgv[0], strings.Join(cmdArgv[1:], " "))
		cmd := exec.Command(cmdArgv[0], cmdArgv[1:]...)

		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		log.Printf("Running command and waiting for it to finish...")
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func Dockerfile(path string, argv []string) error {
	err := Exec(path, []string{
		"docker", "build",
		"-t", "{{.Name}}",
		"-f", "{{.Path}}",
		"{{.ContextDir}}",
	})
	if err != nil {
		return err
	}

	return Exec(path, append([]string{"docker", "run", "-i", "{{.Name}}"}))
}

