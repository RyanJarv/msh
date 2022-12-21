package main

import (
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io"
	"os/exec"
)

func main() {
	p := Process{}
	p.Run()
}

type Provider interface {
	io.ReaderFrom
	io.WriterTo
}

type Process struct {
	stdin  Fd
	stdout Fd
	cmd    *exec.Cmd `json:"-"`
}

func (p Process) Run() error {
	p.stdin = lo.Must(p.cmd.StdinPipe())

	p.stdin

	err := p.cmd.Start()
	if err != nil {
		return err
	}

	line := utils.ReadLine(p.stdout)
	lo.Must(io.WriteString(p.stdin, line))

	return p.cmd.Wait()
}

func (p Process) Stdin() Fd {
	//TODO implement me
	panic("implement me")
}

func (p Process) Stdout() Fd {
	//TODO implement me
	panic("implement me")
}

func (p Process) Config(c Config) Config {
	//TODO implement me
	panic("implement me")
}
