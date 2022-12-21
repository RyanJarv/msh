package fd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io"
	"os"
)

type Config struct {
	WrapperConfig
	AwsCfg         aws.Config
	InitialCommand bool
	StdinSource    io.ReadCloser
	Stdin          io.ReadCloser
	Stdout         io.WriteCloser
	Aggregator     bool
	RemoteProcess  bool
	Name           string
	StdoutArn      string
}

func (c *Config) WriteConf(f *os.File, local bool) {
}

type WrapperConfig struct {
	StdoutUrl   string
	StdoutArn   string
	Pid         *int
	StdoutLocal bool
}

func NewConfig(awsCfg aws.Config, name string) *Config {
	cfg := &Config{
		AwsCfg:        awsCfg,
		Name:          name,
		WrapperConfig: WrapperConfig{},
	}

	if utils.IsTTY(os.Stdin) {
		// Save the pid of the initial process
		cfg.WrapperConfig.Pid = aws.Int(os.Getpid())
		cfg.WrapperConfig.StdoutLocal = true
	} else {
		cfg.WrapperConfig = ReadConf(os.Stdin)
	}

	//if cfg.WrapperConfig.StdoutLocal {
	//	cfg.Stdin = os.Stdin
	//} else {
	//	cfg.Stdin = lo.Must(utils.NewPipe(context.TODO(), cfg.WrapperConfig.Url))
	//}
	//	cfg.WrapperConfig = ReadConf(os.Stdin)
	//	L.Debug.Println("stdin:", cfg.WrapperConfig.Url)
	//
	//	var stdin io.ReadWriteCloser
	//	if !cfg.WrapperConfig.StdoutLocal {
	//		cfg.Stdin = os.Stdin
	//	} else {
	//		cfg.Stdin = lo.Must(utils.NewPipe(context.TODO(), cfg.WrapperConfig.Url))
	//	}
	//
	//	// WrapperConfig command is sending data directly to stdin, forward it to the remote process.
	//	if !cfg.WrapperConfig.StdoutLocal && remote {
	//		go io.Copy(stdin, os.Stdin)
	//	}
	//}

	//if else {
	//	cfg.Arn, cfg.WrapperConfig.Url = SetupSqsFd(cfg.AwsCfg, cfg.Name, "stdout")
	//	cfg.Stdout = lo.Must(utils.NewPipe(context.TODO(), cfg.WrapperConfig.Url))
	//
	//	// Write the updated state to stdout.
	//	WriteConf(os.Stdout, *cfg.WrapperConfig)
	//}
	return cfg
}

func ReadConf(f *os.File) WrapperConfig {
	L.Debug.Println("reading conf from fd", f.Fd())
	scan := bufio.NewScanner(f)
	scan.Scan()

	d := scan.Bytes()
	state := WrapperConfig{}
	err := json.Unmarshal(d, &state)
	if err != nil {
		L.Error.Fatalln("read state:", err)
	}

	L.Debug.Println("got conf:", string(d))
	return state
}

func WriteConf(cfg Config, pipe *OutputPipe) {
	if utils.IsTTY(os.Stdout) {
		return
	}

	state := WrapperConfig{Pid: cfg.WrapperConfig.Pid}
	if pipe != nil {
		state.StdoutArn = pipe.Arn
		state.StdoutUrl = pipe.Url
		state.StdoutLocal = false
	} else {
		state.StdoutLocal = true
	}

	d := lo.Must(json.Marshal(state))
	L.Debug.Println("writing conf:", string(d))
	lo.Must(fmt.Fprint(os.Stdout, string(d)+"\n"))
}
