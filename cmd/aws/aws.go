package aws

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/aws-cdk-go/awscdk/v2/lambdalayerawscli"
	j "github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/samber/lo"
	"log"
	"os"
	"os/exec"
	"strings"
)

//go:embed cli.py
var code []byte

func New(app *app.App, argv []string) (types.CdkStep, error) {
	opts := flag.NewFlagSet("aws", flag.ExitOnError)
	opts.Parse(argv)

	if lo.Contains(flag.Args(), "--help") || lo.Contains(argv, "help") {
		Help(argv)
		os.Exit(1)
	} else if opts.NArg() < 3 {
		return nil, fmt.Errorf("usage: %s <service> <action> <args>", argv)
	}

	iamActions, err := IamActionsFromCliArgs(opts.Arg(1), opts.Arg(2))
	if err != nil {
		return nil, fmt.Errorf("getting iam actions from cli args: %w", err)
	}

	return &AwsCmd{
		IamStatementProps: []awsiam.PolicyStatementProps{
			{
				Actions:   j.Strings(iamActions...),
				Resources: j.Strings("*"),
			},
		},
		Script: string(code),
		Args:   argv,
		Environment: map[string]*string{
			"PYTHONPATH": j.String("/opt/awscli"),
		},
	}, nil
}

func Help(args []string) {
	cmd := exec.Command("aws", args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

type AwsCmd struct {
	Function           awslambda.Function `json:"-"`
	tasks.LambdaInvoke `json:"-"`
	Script             string
	Args               []string
	IamStatementProps  []awsiam.PolicyStatementProps
	Environment        map[string]*string
	IChainable         sfn.IChainable
}

func (s *AwsCmd) GetName() string { return "aws" }

func (s *AwsCmd) BeforeRun(this *types.StepRunInfo) error {
	s.Function = awslambda.NewFunction(this.Scope, this.Id("function"), &awslambda.FunctionProps{
		Runtime:     awslambda.Runtime_PYTHON_3_11(),
		Handler:     j.String("index.lambda_handler"),
		Code:        awslambda.Code_FromInline(j.String(s.Script)),
		Timeout:     awscdk.Duration_Seconds(j.Number(300)),
		Environment: &s.Environment,
	})
	s.Function.AddLayers(lambdalayerawscli.NewAwsCliLayer(this.Scope, this.Id("AwsCliLayer")))

	stmts := lo.Map(s.IamStatementProps, func(item awsiam.PolicyStatementProps, index int) awsiam.PolicyStatement {
		return awsiam.NewPolicyStatement(&item)
	})

	s.Function.Role().AttachInlinePolicy(awsiam.NewPolicy(this.Scope, this.Id("AwsCliPolicy"), &awsiam.PolicyProps{
		Statements: &stmts,
	}))

	return nil
}

func (s *AwsCmd) GetChain(i *types.StepRunInfo) sfn.IChainable {
	return tasks.NewLambdaInvoke(i.Scope, i.Id("invoke"), &tasks.LambdaInvokeProps{
		LambdaFunction: s.Function,
		Payload: sfn.TaskInput_FromObject(&map[string]interface{}{
			"command": sfn.JsonPath_Array(lo.Map(s.Args[1:], func(arg string, index int) *string {
				if strings.HasPrefix(arg, "$") {
					return sfn.JsonPath_StringAt(j.String(arg))
				} else {
					return j.String(arg)
				}
			})...),
		}),
		OutputPath: j.String("$.Payload"),
	}).Next(i.GetChain())
}
