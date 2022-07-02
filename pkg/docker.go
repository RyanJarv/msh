package pkg

import (
	"fmt"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/utils"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	HOME, _  = os.UserHomeDir()
	cacheDir = filepath.Join(HOME, ".msh", "cache")
)

type ExecParams struct {
	Path       string
	ContextDir string
	Name       string
}

func Dockerfile(path string, args ...string) error {
	L.Debug.Printf("executing '%s' with args '%v'\n", path, args)

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
	if !modified(name, path) {
		L.Debug.Printf("unmodified dockerfile, skipping build")
		return nil
	}
	defer updateCache(name)

	dir, err := utils.NewTmpDir(fmt.Sprintf("docker-build-%s", strings.ReplaceAll(path, "/", "-")))
	if err != nil {
		return utils.Wrap(err, "CodeLocal")
	}
	defer func() {
		utils.Report(dir.Remove(), "cleaning up")
	}()

	var out io.Writer
	if os.Getenv("DEBUG") == "" {
		out = utils.DevNull
	}

	return utils.Execute(&utils.ExecuteInput{
		Cmd:    "docker",
		Args:   []string{"build", "-t", name, "-f", path, dir.Path},
		Stdout: out,
		Stderr: out,
	})
}

func updateCache(name string) {
	path := filepath.Join(cacheDir, name)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		err := os.MkdirAll(cacheDir, os.FileMode(0o750))
		if err != nil {
			L.Error.Printf("error creating build cache dir: %s\n", err)
		}

		file, err := os.Create(path)
		if err != nil {
			L.Error.Printf("error creating cache file: %s", err)
		}

		defer file.Close()
	} else {
		t := time.Now().Local()
		err = os.Chtimes(path, t, t)
		if err != nil {
			L.Error.Fatalf("error updating build cache: %s\n", err)
		}
	}
	L.Debug.Printf("updated build cache")
}

func modified(name string, dockerfile string) bool {
	cStat, err := os.Stat(filepath.Join(cacheDir, name))
	if err != nil {
		// Cache doesn't exist, hasn't been built yet.
		return true
	}

	dStat, err := os.Stat(dockerfile)
	if err != nil {
		panic(fmt.Errorf("unable to stat dockerfile: %s", err))
	}

	if dStat.ModTime().After(cStat.ModTime()) {
		return true
	}

	return false
}
