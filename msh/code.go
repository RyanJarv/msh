package msh

import (
	"github.com/ryanjarv/msh/gen"
	"os"
	"path/filepath"
)

func Code(args []string) {
	path := args[0]

	params := ExecParams{
		Path:       path,
		Name:       filepath.Base(path),
	}

	CodeLocal(params, args[1:])
}

func CodeLocal(params ExecParams, args []string) {
	dir := NewTmpDir("code-" + params.Name)
	defer dir.Remove()

	params.ContextDir = dir.Path
	dir.AddFile(params.Path, "app.rb")
	dir.AddContent([]byte(gen.RUBY), "Dockerfile")

	if os.Getenv("DEBUG") == "" {
		TmplExec(params, []string{"docker", "build", "-q", "-t", "{{.Name}}", "{{.ContextDir}}"})
	} else {
		TmplExec(params, []string{"docker", "build", "-t", "{{.Name}}", "{{.ContextDir}}"})
	}
	TmplExec(params, append([]string{"docker", "run", "-p", "127.0.0.1:9000:8080", "-i", "{{.Name}}"}, args...))
}
