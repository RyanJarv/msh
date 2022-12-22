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

func NewCommand(args []string) Command {
	return Command{
		Cmd: exec.Command(args[0], args[1:]...),
		cfg: lo.Must(config.LoadDefaultConfig(context.TODO())),
	}
}

type Command struct {
	*exec.Cmd
	cfg aws.Config
}

func (c Command) Run() error {
	//c.CreatePipeInput.Enrichment = step.Arn()
	L.Debug.Println("Running command", c.Args)
	return c.Cmd.Run()
}
func (c Command) SetStdin(f interface{}) {
	c.Setenv(f)
	c.Cmd.Stdin = f.(io.ReadCloser)
}

func (c Command) GetStdout() io.Reader {
	return lo.Must(c.StdoutPipe())
}

func (c Command) Setenv(f interface{}) {
	switch p := f.(type) {
	case fd.HasEnv:
		c.Env = append(c.Env, p.Env()...)
	}
}
