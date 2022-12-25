package process

import (
	"encoding/json"
	"github.com/ryanjarv/msh/pkg/fd"
	L "github.com/ryanjarv/msh/pkg/logger"
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
	//StderrPipe() (io.ReadCloser, error)
	//StdinPipe() (io.WriteCloser, error)
	//StdoutPipe() (io.ReadCloser, error)
	Run() error
	//Wait()
}

func NewProcess(provider Provider) *Process {
	stdin := fd.NewFd(os.Stdin)
	provider.SetStdin(stdin)

	proc := &Process{
		Stdin:    stdin,
		Provider: provider,
	}

	proc.Stdout = provider.GetStdout()

	if !utils.IsTTY(os.Stdout) || os.Getenv("MSH_STDOUT_TTY") == "false" {
		WriteConf(proc.Stdout)
	}

	return proc
}

type Process struct {
	Provider
	Stdin  interface{}
	Stdout interface{}
}

func (p *Process) Run() error {

	CatchShutdownSignal()

	var wg *sync.WaitGroup

	// If stdout is a TTY, or Stdout of the process can not be passed as a reference, then
	// copy the processes Stdout to the current Stdout.
	if p.Stdout != nil && (utils.IsTTY(os.Stdout) || !fd.IsReferable(p.Stdout)) {
		wg = p.CopyStdout()
	}

	L.Debug.Println("running process")

	err := p.Provider.Run()

	L.Debug.Println("provider run returned", err)

	wg.Wait()

	return err
}

func (p *Process) CopyStdout() *sync.WaitGroup {
	wg := &sync.WaitGroup{}

	if sqs, ok := p.Stdout.(*fd.Sqs); ok {
		L.Debug.Printf("copying: %s -> os.Stdout\n", *sqs.Arn())
	} else {
		L.Debug.Println("copying: p.Stdout -> os.Stdout (%T)", p.Stdout)
	}

	pgid, err := syscall.Getpgid(os.Getpid())
	if err != nil {
		L.Error.Fatalln("failed to get pgid", err)
	}
	L.Debug.Println("pgid:", pgid)

	wg.Add(1)
	go func() {
		fd.MustCopy(os.Stdout, p.Stdout.(io.Reader))
		err := p.Stdout.(io.Closer).Close()
		if err != nil {
			L.Debug.Println("failed to close stdout", err)
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

func WriteConf(out interface{}) {
	d := lo.Must(json.Marshal(out))
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
