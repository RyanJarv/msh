package msh

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/ryanjarv/msh/gen"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const (
	LambdaRiePort   = nat.Port("8080/tcp")
	LocalMaxRuntime = time.Second * 300
)


type Lambda struct {
	Parser
	LambdaExec
	runtime     string
	client      *lambda.Client
	dockerLogin registry.AuthenticateOKBody
}

func NewLambda(cfg Parser) Runnable {
	l := &Lambda{
		Parser:    cfg,
		runtime:   cfg.Args[0],
		LambdaExec: NewLambdaExec(cfg.Global),
		client:     lambda.NewFromConfig(cfg.Aws.Config),
	}
	l.Path = cfg.Args[1]
	return l
}

func (l Lambda) Run() (err error) {
	switch l.runtime {
	case "ruby":
		l.LambdaRuby()
	default:
		panic(errors.New("lambda runtime not supported: " + l.runtime))
	}

	if os.Getenv("MSH_CONTEXT_TRIGGER_DEPLOY") == "true" {
		l.Deploy()
	} else {
		l.Local()
	}
	return
}

func (l Lambda) Local() {
	l.Start()
	l.StdinToHttp(url.URL{
		Scheme: "http",
		Host: 	fmt.Sprintf("127.0.0.1:%s", l.Inspect.NetworkSettings.Ports[LambdaRiePort][0].HostPort),
		Path:   "/2015-03-31/functions/function/invocations",
	})

	go l.Logs()

	go func() {
		time.Sleep(LocalMaxRuntime)
		l.Stop()
	}()
	l.Wait()
}

func (l *Lambda) Deploy() {
	host := fmt.Sprintf("%s.dkr.ecr.%s.amazonaws.com", *l.Aws.Identity.Account, l.Aws.Config.Region)
	repoPath := path.Join("msh", l.Project)
	l.remoteUri = fmt.Sprintf("%s:%s", path.Join(host, repoPath), l.Tag)

	l.CreateRepo(repoPath)

	l.NewTag(l.remoteUri)

	l.Push(*ecr.NewFromConfig(l.Aws.Config), l.remoteUri)

	g := gen.SAM

	cfn := NewCfn(Parser{
		Content: g,
		Context: l.Context,
		Global:  l.Global,
	})
	cfn.Parse(map[string]string{
		"ImageUri": l.remoteUri,
		"FunctionName": strings.ToTitle(strings.TrimSuffix(filepath.Base(l.Path), ".rb")),
	})

	err := cfn.Run()
	if err != nil {
		panic(err)
	}

	var endpoint string
	for _, v := range cfn.(*Cfn).StackOutputs {
		if *v.OutputKey == "APPApi" {
			endpoint = *v.OutputValue
		}
	}
	parse, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	l.StdinToHttp(*parse)
	time.Sleep(time.Hour * 100) // FIXME: asdf
}


func (l *Lambda) LambdaRuby() {
	script, err := ioutil.ReadFile(l.Path)
	if err != nil {
		panic(err)
	}

	tarGz := MakeTarGz(&TarFiles{
		{Name: filepath.Base(l.Path), Body: string(script)},
		{Name: "Dockerfile", 		Body: gen.RUBY},
	})

	l.Build(tarGz)
}

