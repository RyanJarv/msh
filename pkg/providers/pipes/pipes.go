package pipes

import (
	pipesTypes "github.com/aws/aws-sdk-go-v2/service/pipes/types"
	"github.com/ryanjarv/msh/pkg/fd"
	"github.com/samber/lo"
	"io"
	"os"
	"sync"
)

func Connect(previous, current *Process, aggregate bool) {
	go lo.Must(current.stdin.ReadFrom(previous.Stdout()))
}

type AwsPipe struct {
	pipesTypes.PipeSourceParameters
	pipesTypes.PipeTargetParameters
}

func (s AwsPipe) NewFd() fd.Sqs {
	panic("not implemented")
}

func (s AwsPipe) ReadFrom(r interface{}) (n int64, err error) {
	switch source := r.(type) {
	case pipesTypes.PipeSourceParameters:
		// Just read directly from the sqs output queue of the last command.
		s.PipeSourceParameters = source

		// os.Stdin will close to signal when processing has completed.
		return io.Copy(os.Stdout, os.Stdin)
	case io.ReadCloser:
		sqs := s.NewFd()
		s.PipeSourceParameters = sqs.PipeSourceParameters
		return io.Copy(sqs, source)
	default:
		panic("not implemented")
	}
}

func (s AwsPipe) WriteTo(r interface{}) (n int64, err error) {
	switch target := r.(type) {
	case pipesTypes.PipeTargetParameters:
		// Just read directly from the sqs output queue of the last command.
		s.PipeTargetParameters = target

		// TODO: return when stdin returns.
		Block()
		return 0, nil
	case io.Writer:
		sqs := s.NewFd()
		s.PipeTargetParameters = sqs.PipeTargetParameters
		return io.Copy(target, sqs)
	default:
		panic("not implemented")
	}
}

type LocalPipe struct {
	io.ReadCloser
	Ref string
}

func Block() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
