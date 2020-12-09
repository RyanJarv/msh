package main

import (
	"context"
	"flag"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/pkg/errors"
	L "github.com/ryanjarv/msh/logger"
	"github.com/ryanjarv/msh/msh"
	"log"
	"os"
)

//go:generate go run script/templates.go

func main() {
	if err := run(); err != nil {
		if err, ok := err.(msh.StackTracer); ok {
			for _, f := range err.StackTrace() {
				L.Error.Printf("%+s:%d\n", f, f)
			}
		}
		log.Fatalln(err)
	}
}

var args = flag.CommandLine

func run() (err error) {
	if err = args.Parse(os.Args[1:]); err != nil {
		return errors.Wrap(err, "mode arg missing")
	}

	L.Debug.Println("args: %v", args.Args())

	cfg := NewConfig()

	switch cmd := args.Args()[0]; cmd {
	// {
	//		"config": {
	//			"build": {}
	// 		}
	// }
	//case "build":
	//	err = msh.Build(args.Args())
	case "lambda":
		err = msh.NewLambda(cfg).Run()
		//err = msh.NewCfn(cfg).Parse().Run()
	case "compose":
		err = msh.Compose(args.Args()[1:])
	case "dockerfile":
		err = msh.Dockerfile(args.Args()[1:])
	case "deploy":
		err = msh.Deploy(args.Args()[1:])
	case "cfn":
		err = msh.NewCfn(cfg).Run()
	//case "init":
	//	err = msh.Init(args.Args()[1:])
	case "":
		args.Usage()
		err = errors.New("Must specify subcommand")
	default:
		return errors.Errorf("unknown command: '%v'", cmd)
	}
	return err
}

func NewConfig() msh.Parser {
	cfg := msh.Parser{
		Context: context.Background(),
		Args:    args.Args()[1:],
		Global: msh.Global{
			Project: "msh-default",
		},
	}
	var err error
	if cfg.Global.Aws.Config, err = config.LoadDefaultConfig(config.WithDefaultRegion("us-east-1")); err != nil {
		panic(err)
	}

	client := sts.NewFromConfig(cfg.Aws.Config)
	identity, err := client.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	if err != nil {
		panic(err)
	}
	cfg.Global.Aws.Identity = *identity
	return cfg
}

