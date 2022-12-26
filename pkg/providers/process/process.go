package process

import (
	"encoding/json"
	"github.com/ryanjarv/msh/pkg/fd"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Provider interface {
	SetStdin(interface{})
	GetStdout() io.ReadCloser
	Run() error
}

func NewProcess(provider Provider) *Process {
	stdin, localStdin := fd.NewFd(os.Stdin)
	provider.SetStdin(stdin)

	proc := &Process{
		Stdin:      stdin,
		Provider:   provider,
		LocalStdin: localStdin,
	}

	proc.Stdout = provider.GetStdout()
	if !utils.IsTTY(os.Stdout) || os.Getenv("MSH_STDOUT_TTY") == "false" {
		// Write a reference to stdout so the next command can determine where to find the input.
		WriteConf(proc.Stdout.(io.ReadCloser))
	}

	return proc
}

type Process struct {
	Provider
	Stdin  interface{}
	Stdout interface{}

	// LocalStdin is set to true when stdin is local.
	LocalStdin bool
}

func (p *Process) Deploy() error {
	if d, ok := p.Provider.(types.Deployable); ok {
		return utils.Wrap(d.Deploy(), "deploying %T", d)
	}
	return nil
}

func (p *Process) Run() error {
	CatchShutdownSignal()

	var wg *sync.WaitGroup

	// If Stdout is a tty or can not be passed as a reference, copy locally.
	if p.Stdout != nil && (utils.IsTTY(os.Stdout) || !fd.IsReferable(p.Stdout)) {
		wg = p.CopyStdout()
	}

	L.Debug.Println("running process")

	err := p.Provider.Run()

	L.Debug.Println("provider run returned", err)

	if wg != nil {
		L.Debug.Println("waiting for stdout copy to finish")
		wg.Wait()
	}

	return err
}

func (p *Process) CopyStdout() *sync.WaitGroup {
	wg := &sync.WaitGroup{}

	if sqs, ok := p.Stdout.(*fd.Sqs); ok {
		L.Debug.Printf("copying: %s -> os.Stdout\n", *sqs.Arn())
	} else {
		L.Debug.Printf("copying: p.Stdout -> os.Stdout (%T)\n", p.Stdout)
	}

	pgid, err := syscall.Getpgid(os.Getpid())
	if err != nil {
		L.Error.Fatalln("failed to get pgid", err)
	}
	L.Debug.Println("pgid:", pgid)

	wg.Add(1)
	go func() {
		fd.MustCopy(os.Stdout, p.Stdout.(io.Reader))

		// We don't monitor remote processes, so unless we are receiving input locally, we won't know when to close
		// stdout. To work around this we send the closed event once per pipeline and forward it to the next process
		// remotely. This means we need to avoid closing a remote pipeline more than once across separate commands.
		//
		// Remote pipelines don't correlate to local pipelines since a non-msh command in the middle will cause data
		// to be proxied through it, instead only close stdout if stdin was not a reference.
		sqs, ok := p.Stdout.(*fd.Sqs)
		if ok && p.LocalStdin && utils.IsTTY(os.Stdin) {
			err := sqs.Close()
			if err != nil {
				L.Debug.Println("failed to close stdout", err)
			}
		}

		if r, ok := p.Stdout.(io.Closer); ok {
			err = r.Close()
			if err != nil {
				L.Debug.Println("failed to close stdin", err)
			}
		}

		err = os.Stdin.Close()
		if err != nil {
			L.Debug.Println("failed to close stdin", err)
		}
		//err = syscall.Kill(-pgid, syscall.SIGUSR1)
		//if err != nil {
		//	L.Error.Println("failed to kill pgid", err)
		//}

		wg.Done()
	}()
	return wg
}

func WriteConf(stdout io.ReadCloser) {
	d := lo.Must(json.Marshal(stdout))
	lo.Must(io.WriteString(os.Stdout, string(d)+"\n"))
}

// CatchShutdownSignal catches the process group shutdown signal and exits.
//
// SIGPIPE doesn't appear to be catchable, so we use SIGUSR1 instead.
func CatchShutdownSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			L.Debug.Println("caught SIGUSR1, shutting down")
			os.Stdout.Sync()
			os.Stdout.Sync()
			os.Exit(0)
		}
	}()
}
