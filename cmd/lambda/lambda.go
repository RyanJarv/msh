package lambda

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
	"os"
)

type LambdaOpts struct {
	InputPath      *string
	ResultSelector *map[string]interface{}
	OutputPath     *string
	ResultPath     *string
}

func New(app app.App) (*Lambda, error) {
	path := app.Arg(1)

	env := map[string]*string{}
	for i, arg := range flag.Args()[2:] {
		env[fmt.Sprintf("ARG%d", i+1)] = &arg

	}

	script, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	if !bytes.Contains(script, []byte("lambda_handler")) {
		return nil, fmt.Errorf("script must contain a `lambda_handler` function")
	}

	return &Lambda{
		Script:         string(script),
		Args:           app.Args(),
		Environment:    env,
		TimeoutSeconds: 300,
	}, nil
}

type Lambda struct {
	*LambdaOpts
	types.IChain
	awslambda.Function
	Script         string
	Args           []string
	Environment    map[string]*string
	TimeoutSeconds int
}

func (s Lambda) GetName() string { return "lambda" }

func (s *Lambda) Compile(stack constructs.Construct, i int) error {
	s.Function = awslambda.NewFunction(stack, jsii.String(s.GetName()), &awslambda.FunctionProps{
		Runtime:     awslambda.Runtime_PYTHON_3_11(),
		Handler:     jsii.String("index.lambda_handler"),
		Code:        awslambda.Code_FromInline(jsii.String(s.Script)),
		Environment: &s.Environment,
		Timeout:     awscdk.Duration_Seconds(jsii.Number(s.TimeoutSeconds)),
	})

	s.IChain = tasks.NewLambdaInvoke(stack, jsii.String(fmt.Sprintf("%s-invoke", s.GetName())), &tasks.LambdaInvokeProps{
		LambdaFunction: s.Function,
		InputPath:      s.InputPath,
		ResultSelector: s.ResultSelector,
		OutputPath:     s.OutputPath,
		ResultPath:     s.ResultPath,
	})

	return nil
}

func (s *Lambda) SetOpts(opts *LambdaOpts) {
	s.LambdaOpts = opts
}
