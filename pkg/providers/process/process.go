package process

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/ryanjarv/msh/pkg/fd"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Provider interface {
	SetStdin(interface{})
	GetStdout() io.Reader
	Run() error
}

func NewProcess(provider Provider) *Process {
	proc := &Process{
		Provider: provider,
	}

	if utils.IsTTY(os.Stdin) {
		// Save the pid of the initial process
		proc.Pid = aws.Int(os.Getpid())
	} else {
		proc.ReadConf(os.Stdin)
	}

	proc.SetStdin(proc.Stdin)
	proc.Stdout = proc.GetStdout()

	if !utils.IsTTY(os.Stdout) || os.Getenv("MSH_STDOUT_TTY") == "false" {
		proc.WriteConf(os.Stdout)
	}

	return proc
}

type Process struct {
	Provider Provider    `json:"-"`
	Stdin    interface{} `json:",omitempty"`
	Stdout   interface{} `json:"-"`
	Pid      *int        `json:"-"`
}

func (p *Process) Run() error {
	wg := sync.WaitGroup{}

	CatchShutdownSignal()

	// If stdout is a TTY, or Stdout of the process can not be passed as a reference, then
	// copy the processes Stdout to the current Stdout.
	if utils.IsTTY(os.Stdout) || !fd.IsReferable(p.Stdout) {
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

			L.Debug.Println("pgid is %d, sending SIGKILL", pgid)
			err = syscall.Kill(-pgid, syscall.SIGUSR1)
			if err != nil {
				L.Error.Println("failed to kill pgid", err)
			}

			wg.Done()
		}()
	}

	err := p.Provider.Run()
	L.Debug.Println("provider run returned", err)

	wg.Wait()
	L.Debug.Println("all done")
	return err
}

func (p *Process) SetStdin(stdin interface{}) {
	if utils.IsNilOrEmpty(stdin) {
		L.Debug.Println("stdin set to: os.Stdin")
		stdin = os.Stdin
	}
	L.Debug.Printf("stdin (%T) set to: %+v", stdin, stdin)

	switch s := stdin.(type) {
	case map[string]interface{}:
		url := s["Url"].(string)
		L.Debug.Println("stdin is sqs, skipping copy:", url)
		stdin = lo.Must(fd.OpenSqs(context.TODO(), url))
	}

	p.Provider.SetStdin(stdin)
}

func (p *Process) GetStdout() io.Reader {
	return p.Provider.GetStdout()
}

func (p *Process) WriteConf(w *os.File) {
	d := lo.Must(json.Marshal(Process{
		Stdin: p.Stdout,
	}))
	lo.Must(io.WriteString(w, string(d)+"\n"))
	time.Sleep(100 * time.Millisecond)
}

func (p *Process) ReadConf(f *os.File) {
	L.Debug.Println("reading conf from fd", f.Fd())

	d := lo.Must(readLine(f))

	L.Debug.Println("config json:", string(d))

	err := json.Unmarshal(d, p)
	if err != nil {
		L.Error.Fatalln("failed to read config from stdin", err)
	}

	L.Debug.Println("got conf:", string(d))
}

func readLine(f *os.File) ([]byte, error) {
	b := make([]byte, 1, 1)
	var s []byte
	for {
		_, err := f.Read(b)
		if err != nil {
			return nil, fmt.Errorf("read byte: %w", err)
		}
		if b[0] == '\n' {
			break
		}

		s = append(s, b...)

	}

	return s, nil
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
