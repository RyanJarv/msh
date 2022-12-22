package process

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/ryanjarv/msh/pkg/fd"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io"
	"os"
	"sync"
	"time"
)

type Provider interface {
	SetStdin(interface{})
	GetStdout() io.Reader
	Run() error
}

//type Command struct {
//	stdin Fd
//	stdout Fd
//	cmd    *exec.Cmd `json:"-"`
//}

func NewProcess(provider Provider) *Process {
	proc := &Process{
		Provider: provider,
		//OsStdin:  os.Stdin,
	}

	if utils.IsTTY(os.Stdin) && os.Getenv("MSH_STDIN_TTY") != "false" {
		// Save the pid of the initial process
		proc.Pid = aws.Int(os.Getpid())
	} else {
		proc.ReadConf(os.Stdin)
	}

	//proc.Stdin = proc.SetStdin(proc.Stdin)
	proc.SetStdin(proc.Stdin)
	proc.Stdout = proc.GetStdout()

	if !utils.IsTTY(os.Stdout) || os.Getenv("MSH_STDOUT_TTY") == "false" {
		proc.WriteConf(os.Stdout)
	}

	return proc
}

type Process struct {
	Provider Provider `json:"-"`
	OsStdin  io.Reader
	Stdin    interface{} `json:",omitempty"`
	Stdout   interface{} `json:"-"`
	Pid      *int        `json:"-"`
}

func (p *Process) Run() error {
	wg := sync.WaitGroup{}

	if utils.IsTTY(os.Stdout) || !fd.IsReferable(p.Stdout) {
		L.Debug.Println("copying: p.Stdout -> os.Stdout")
		wg.Add(1)
		go func() {
			fd.MustCopy(os.Stdout, p.Stdout.(io.Reader))
			wg.Done()
		}()
	}

	err := p.Provider.Run()
	L.Debug.Println("provider run returned", err)

	wg.Wait()
	return err
}

func (p *Process) SetStdin(s interface{}) {
	if utils.IsEmptyOrEmpty(s) {
		L.Debug.Println("stdin set to: os.Stdin")
		s = os.Stdin
	}

	p.Provider.SetStdin(s)
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
