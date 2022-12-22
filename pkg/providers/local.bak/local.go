package local

import (
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/ryanjarv/msh/pkg/conf"
	"github.com/ryanjarv/msh/pkg/fd"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

import (
	"context"
)

// Run takes the invocation args and transparently runs the intended command while intercepting sensitive info.
func Run(args []string) error {
	ourPath := args[0]
	cmdPath, err := LookPathSkip(filepath.Base(ourPath), ourPath)
	if err != nil {
		return err
	}

	cmd := exec.Command(cmdPath, args[1:]...)
	L.Debug.Println("running command:", cmdPath, strings.Join(args[1:], " "))

	name := utils.GetName(args)
	awsCfg := lo.Must(config.LoadDefaultConfig(context.TODO()))

	cfg := conf.NewConfig(awsCfg, name)
	conf.WriteConf(*cfg, nil)

	cmd.Stdin = fd.NewInputPipe(cfg, false)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cfg.WriteConf(os.Stdout, true)

	return cmd.Run()
}

type Config struct{}

func NewProcess(args []string) Process {
	cmd := exec.Command(args[0], args[1:]...)

	return Process{
		cmd: cmd,
	}
}

func findExecutable(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return nil
	}
	return fs.ErrPermission
}

// LookPathSkip is the same as exec/LookPath but skips the given skipPath during lookup.
func LookPathSkip(file, skipPath string) (string, error) {
	path := os.Getenv("PATH")
	for _, dir := range filepath.SplitList(path) {
		if filepath.Clean(dir) == filepath.Clean(skipPath) {
			continue
		}
		if dir == "" {
			// Unix shell semantics: path element "" means "."
			dir = "."
		}
		path := filepath.Join(dir, file)
		if err := findExecutable(path); err == nil {
			return path, nil
		}
	}
	return "", &exec.Error{Name: file, Err: exec.ErrNotFound}
}

//type Wrapper struct {
//	previous *Process
//}
//
//func (w Wrapper) Previous() *Process {
//	if utils.IsTTY(os.Stdin) {
//		return nil
//	}
//
//	if w.previous == nil {
//		line := ReadLine(os.Stdin)
//		w.previous = &Process{}
//		err := json.Unmarshal([]byte(line), w.previous)
//		if err != nil {
//			L.Error.Fatalln("failed reading previous process config")
//		}
//	}
//
//	return w.previous
//}
//
