package pkg

import (
	"fmt"
	L "github.com/ryanjarv/msh/logger"
	"github.com/ryanjarv/msh/pkg/utils"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ExecParams struct {
	Path       string
	ContextDir string
	Name       string
}

func Dockerfile(path string, args ...string) error {
	L.Debug.Printf("Running dockerExec with path '%s' and args '%v'\n", path, args)

	name := filepath.Base(path)

	err := build(name, path)
	if err != nil {
		return utils.Wrap(err, "Dockerfile")
	}

	return run(name, args)
}

func run(name string, args []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return utils.Wrap(err, "dockerExec")
	}

	return utils.Execute(&utils.ExecuteInput{
		Cmd:  "docker",
		Args: append([]string{"run", "-v", fmt.Sprintf("%s:/app", pwd), "-it", name}, args...),
	})
}

func build(name, path string) error {
	dir, err := utils.NewTmpDir(fmt.Sprintf("docker.build.%s", strings.ReplaceAll(path, "/", "-")))
	if err != nil {
		return utils.Wrap(err, "CodeLocal")
	}
	defer utils.Report(dir.Remove(), "cleaning up")

	var out io.Writer
	if os.Getenv("DEBUG") == "" {
		out = utils.DevNull
	}

	return utils.Execute(&utils.ExecuteInput{
		Cmd:    "docker",
		Args:   []string{"run", "-t", name, "-f", path, dir.Path},
		Stdout: out,
	})
}
