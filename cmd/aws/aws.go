package aws

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/aws-cdk-go/awscdk/v2/lambdalayerawscli"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/samber/lo"
	"os"
	"os/exec"
	"strings"
)

//go:embed cli.py
var code []byte

func New(app app.App) (*AwsCmd, error) {
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
				Actions:   jsii.Strings(iamActions...),
				Resources: jsii.Strings("*"),
			},
		},
		Script: string(code),
		Args:   app.OsArgs,
		Environment: map[string]*string{
			"PYTHONPATH": jsii.String("/opt/awscli"),
		},
	}, nil
}

func Help(args []string) {
	cmd := exec.Command("aws", args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type AwsCmd struct {
	awslambda.Function `json:"-"`
	tasks.LambdaInvoke
	Script            string
	Args              []string
	IamStatementProps []awsiam.PolicyStatementProps
	Environment       map[string]*string
}

func (s AwsCmd) GetName() string      { return "aws" }
func (s AwsCmd) getName(i int) string { return fmt.Sprintf("aws-%d", i) }

func (s *AwsCmd) Compile(stack constructs.Construct, i int) error {
	s.Function = awslambda.NewFunction(stack, jsii.String(s.getName(i)), &awslambda.FunctionProps{
		Runtime:     awslambda.Runtime_PYTHON_3_11(),
		Handler:     jsii.String("index.lambda_handler"),
		Code:        awslambda.Code_FromInline(jsii.String(s.Script)),
		Timeout:     awscdk.Duration_Seconds(jsii.Number(300)),
		Environment: &s.Environment,
	})
	s.Function.AddLayers(lambdalayerawscli.NewAwsCliLayer(stack, jsii.String("AwsCliLayer")))

	stmts := lo.Map(s.IamStatementProps, func(item awsiam.PolicyStatementProps, index int) awsiam.PolicyStatement {
		return awsiam.NewPolicyStatement(&item)
	})

	s.Function.Role().AttachInlinePolicy(awsiam.NewPolicy(stack, jsii.String("AwsCliPolicy"), &awsiam.PolicyProps{
		Statements: &stmts,
	}))

	args := lo.Map(s.Args[1:], func(arg string, index int) *string {
		if strings.HasPrefix(arg, "$") {
			return awsstepfunctions.JsonPath_StringAt(jsii.String(arg))
		} else {
			return jsii.String(arg)
		}
	})

	s.LambdaInvoke = tasks.NewLambdaInvoke(stack, jsii.String(fmt.Sprintf("%s-invoke", s.getName(i))), &tasks.LambdaInvokeProps{
		LambdaFunction: s.Function,
		Payload: awsstepfunctions.TaskInput_FromObject(&map[string]interface{}{
			"command": awsstepfunctions.JsonPath_Array(args...),
		}),
		OutputPath: jsii.String("$.Payload"),
	})

	return nil
}
