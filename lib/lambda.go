package lib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	ecrTypes "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/ryanjarv/msh/gen"
	L "github.com/ryanjarv/msh/logger"
	"io"
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

	runtime     string
	client      *lambda.Client
	ecrClient   *ecr.Client
	docker      LambdaExec
	dockerLogin registry.AuthenticateOKBody
}

func NewLambda(cfg Parser) Runnable {
	l := &Lambda{
		Parser:    cfg,
		runtime:   cfg.Args[0],
		client:    lambda.NewFromConfig(cfg.Aws.Config),
		ecrClient: ecr.NewFromConfig(cfg.Aws.Config),
	}
	l.Path = cfg.Args[1]
	return l
}

func (l Lambda) Run() (err error) {
	switch l.runtime {
	case "ruby":
		l.docker = LambdaRuby(l.Path, l.Project)
	default:
		panic(errors.New("lambda runtime not supported: " + l.runtime))
	}

	if os.Getenv("MSH_CONTEXT_TRIGGER_DEPLOY") == "true" {
		l.Deploy(l.docker)
	} else {
		l.Local(l.docker)
	}
	return
}

func (l Lambda) Local(cmd LambdaExec) {
	cmd.ContainerStart()

	cmd.StdinToHttp(url.URL{
		Scheme: "http",
		Host: 	fmt.Sprintf("127.0.0.1:%s", cmd.PortBinding.HostPort),
		Path:   "/2015-03-31/functions/function/invocations",
	})

	go cmd.ContainerLogs()
	go func() {
		time.Sleep(LocalMaxRuntime)
		cmd.ContainerStop()
	}()
	cmd.ContainerWait()
}

func (l Lambda) Deploy(cmd LambdaExec) {
	host := fmt.Sprintf("%s.dkr.ecr.%s.amazonaws.com", *l.Aws.Identity.Account, l.Aws.Config.Region)
	repoPath := path.Join("msh", l.Project)
	l.docker.remoteUri = fmt.Sprintf("%s:%s", path.Join(host, repoPath), l.docker.localTag)

	l.createRepo(repoPath)

	if err := l.docker.cli.ImageTag(l.Context, l.docker.localTag, l.docker.remoteUri); err != nil {
		panic(err)
	}

	l.pushImage()

	g := gen.SAM

	cfn := NewCfn(Parser{
		Content: g,
		Context: l.Context,
		Global:  l.Global,
	})
	cfn.Parse(map[string]string{
		"ImageUri": l.docker.remoteUri,
		"FunctionName": strings.ToTitle(strings.TrimSuffix(filepath.Base(l.Path), ".rb")),
	})

	err := cfn.Run()
	if err != nil {
		panic(err)
	}

	var endpoint string
	for _, v := range cfn.(*Cfn).StackOutputs {
		fmt.Println(v.ExportName)
		if *v.OutputKey == "APPApi" {
			endpoint = *v.OutputValue
		}
	}
	parse, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	cmd.StdinToHttp(*parse)
	time.Sleep(time.Hour * 100) // FIXME: asdf
}

func (l Lambda) createRepo(repoPath string) {
	_, err := l.ecrClient.CreateRepository(l.Context, &ecr.CreateRepositoryInput{RepositoryName: aws.String(repoPath)})
	var exists *ecrTypes.RepositoryAlreadyExistsException
	if err != nil && !errors.As(err, &exists) {
		panic(err)
	}

}

func (l Lambda) pushImage() {
	token, err := l.ecrClient.GetAuthorizationToken(l.Context, &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		panic(err)
	}

	// SDK seems to encode it in the wrong format so, decode and reencode
	auth := types.AuthConfig{
		Username: "AWS",
	} // UserName is always AWS
	if s, err := base64.StdEncoding.DecodeString(*token.AuthorizationData[0].AuthorizationToken); err != nil {
		panic(err)
	} else {
		auth.Password = strings.Split(string(s), ":")[1] // user:pass
	}
	encodedJSON, err := json.Marshal(auth)
	if err != nil {
		panic(err)
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	push, err := l.docker.cli.ImagePush(l.Context, l.docker.remoteUri, types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		panic(err)
	}
	defer push.Close()

	_, err = io.Copy(L.Debug.Writer(), push)
}

func LambdaRuby(path, name string) LambdaExec {
	script, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	tarGz := MakeTarGz(&TarFiles{
		{Name: filepath.Base(path), Body: string(script)},
		{Name: "Dockerfile", 		Body: gen.RUBY},
	})

	lambda := NewLambdaExec(GetRandStr())
	lambda.BuildContainer(tarGz)

	return lambda
}

