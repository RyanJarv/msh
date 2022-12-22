package main

import (
	"context"
	"flag"
	"github.com/ryanjarv/msh/pkg/fd"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/process"
	"github.com/samber/lo"
	"io"
)

func main() {
	flag.Parse()

	sqs := &Sqs{}

	proc := process.NewMsh(sqs)
	err := proc.Run()

	if err != nil {
		L.Error.Fatalln(err)
	}
}

type Sqs struct {
	*fd.Sqs
}

func (s *Sqs) Run() error {
	return nil
}

func (s *Sqs) SetStdin(p interface{}) {
	s.Sqs = lo.Must(fd.NewSqsFrom(context.TODO(), p))
}

func (s *Sqs) GetStdout() io.Reader {
	return s.Sqs
}
