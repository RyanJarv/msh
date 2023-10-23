package command

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/ryanjarv/msh/pkg/fd"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/samber/lo"
	"io"
	"os/exec"
)

type ICommand interface {
	SetStdin(interface{})
	GetStdout() io.ReadCloser
}

func NewCommand(args []string) Command {
	if len(args) == 0 {
		L.Error.Fatalln("no command provided")
	}

	return Command{
		Cmd: exec.Command(args[0], args[1:]...),
		cfg: lo.Must(config.LoadDefaultConfig(context.TODO())),
	}
}

type Command struct {
	*exec.Cmd
	cfg aws.Config
}

func (c Command) SetStdin(f interface{}) {
	c.setenv(f)
	pipe, err := c.Cmd.StdinPipe()
	if err != nil {
		L.Error.Fatalln("failed to get stdin pipe", err)
	}

	go fd.MustCopy(pipe, f.(io.Reader))
}

func (c Command) GetStdout() io.ReadCloser {
	return lo.Must(c.StdoutPipe())
}

func (c Command) setenv(f interface{}) {
	switch p := f.(type) {
	case fd.HasEnv:
		c.Env = append(c.Env, p.Env()...)
	}
}
