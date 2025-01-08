package aws

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/aws-cdk-go/awscdk/v2/lambdalayerawscli"
	"github.com/aws/constructs-go/constructs/v10"
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

func New(app *app.App) (*AwsCmd, error) {
	if lo.Contains(app.OsArgs, "--help") || lo.Contains(app.OsArgs, "help") {
		Help(app.OsArgs)
		os.Exit(1)
	}

	iamActions, err := IamActionsFromCliArgs(app.OsArgs)
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
		Args:   app.OsArgs,
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
	awslambda.Function `json:"-"`
	tasks.LambdaInvoke `json:"-"`
	Script             string
	Args               []string
	IamStatementProps  []awsiam.PolicyStatementProps
	Environment        map[string]*string
	IChainable         sfn.IChainable
}

func (s *AwsCmd) GetName() string { return "aws" }

func (s *AwsCmd) Init(scope constructs.Construct, last, next types.CdkStep, i int) error {
	name := fmt.Sprintf("%s-%d", s.GetName(), i)
	s.Function = awslambda.NewFunction(scope, j.String(name+"-function"), &awslambda.FunctionProps{
		Runtime:     awslambda.Runtime_PYTHON_3_11(),
		Handler:     j.String("index.lambda_handler"),
		Code:        awslambda.Code_FromInline(j.String(s.Script)),
		Timeout:     awscdk.Duration_Seconds(j.Number(300)),
		Environment: &s.Environment,
	})
	s.Function.AddLayers(lambdalayerawscli.NewAwsCliLayer(scope, j.String("AwsCliLayer")))

	stmts := lo.Map(s.IamStatementProps, func(item awsiam.PolicyStatementProps, index int) awsiam.PolicyStatement {
		return awsiam.NewPolicyStatement(&item)
	})

	s.Function.Role().AttachInlinePolicy(awsiam.NewPolicy(scope, j.String("AwsCliPolicy"), &awsiam.PolicyProps{
		Statements: &stmts,
	}))

	return nil
}

func (s *AwsCmd) GetChain(step *types.StepRunInfo) sfn.IChainable {
	invoke := tasks.NewLambdaInvoke(step.Scope, step.Id("invoke"), &tasks.LambdaInvokeProps{
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
	})

	switch v := step.Next.Step.(type) {
	case types.SfnChainable:
		return sfn.Chain_Sequence(invoke, v.GetChain(step.Next))
	default:
		return invoke
	}
}
