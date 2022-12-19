package local

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/ryanjarv/msh/pkg/utils/fd"
	"github.com/samber/lo"
	"io"
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

	cfg := lo.Must(config.LoadDefaultConfig(context.TODO()))
	name := utils.NameFromCmd(args)

	stdin, stdout := setupFds(cfg, name)
	defer stdout.Close()

	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func setupFds(cfg aws.Config, name string) (stdin io.ReadCloser, stdout io.WriteCloser) {
	var conf *fd.StateJson
	if !utils.IsTTY(os.Stdin) {
		conf = fd.ReadConf(os.Stdin)
	}

	if conf == nil {
		L.Debug.Println("stdin: terminal")
		stdin = os.Stdin
	} else {
		L.Debug.Println("stdin:", conf.StdinUrl)
		stdin = lo.Must(utils.NewPipe(context.TODO(), conf.StdinUrl))
	}

	if utils.IsTTY(os.Stdout) {
		L.Debug.Println("stdout: terminal")
		stdout = os.Stdout
	} else {
		arn, url := fd.SetupSqsFd(cfg, name, "stdout")
		L.Debug.Println("stdout:", arn)
		stdout = lo.Must(utils.NewPipe(context.TODO(), url))

		fd.WriteConf(os.Stdout, fd.StateJson{
			StdinArn: arn,
			StdinUrl: url,
			Pid:      aws.Int(os.Getpid()),
		})
	}

	return stdin, stdout
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
	return "", &exec.Error{file, exec.ErrNotFound}
}
