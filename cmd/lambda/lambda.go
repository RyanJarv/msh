package lambda

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
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

func New(app *app.App) (*Lambda, error) {
	return NewInternal(app.Flag.Args()[min(app.Flag.NArg(), 1):])
}

// NewInternal can be called directly from other commands that use a lambda without needing to parse the app.
func NewInternal(args []string) (*Lambda, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("lambda: must provide a path to a python script, got: %s", args)
	}
	path := args[0]

	env := map[string]*string{}
	for i, arg := range args[1:] {
		env[fmt.Sprintf("ARG%d", i)] = jsii.String(arg)

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
		Args:           args,
		Environment:    env,
		TimeoutSeconds: 300,
		LambdaOpts:     &LambdaOpts{},
	}, nil
}

type Lambda struct {
	sfn.IChainable `json:"-"`
	function       awslambda.Function
	LambdaOpts     *LambdaOpts
	Script         string
	Args           []string
	Environment    map[string]*string
	TimeoutSeconds int
}

func (s Lambda) GetName() string { return "lambda" }

func (s *Lambda) Init(scope constructs.Construct, i int) error {
	name := fmt.Sprintf("%s-%d", s.GetName(), i)
	s.function = awslambda.NewFunction(scope, jsii.String(name), &awslambda.FunctionProps{
		Runtime:     awslambda.Runtime_PYTHON_3_11(),
		Handler:     jsii.String("index.lambda_handler"),
		Code:        awslambda.Code_FromInline(jsii.String(s.Script)),
		Environment: &s.Environment,
		Timeout:     awscdk.Duration_Seconds(jsii.Number(s.TimeoutSeconds)),
	})

	s.IChainable = tasks.NewLambdaInvoke(scope, jsii.String(fmt.Sprintf("%s-invoke", s.GetName())), &tasks.LambdaInvokeProps{
		LambdaFunction: s.function,
		InputPath:      s.LambdaOpts.InputPath,
		ResultSelector: s.LambdaOpts.ResultSelector,
		OutputPath:     s.LambdaOpts.OutputPath,
		ResultPath:     s.LambdaOpts.ResultPath,
	})

	return nil
}

func (s *Lambda) GetChain(step *types.StepRunInfo) (err error) {
	switch v := step.Next.Step.(type) {
	case sfn.IChainable:
		s.IChainable = sfn.Chain_Sequence(v.StartState(), s.IChainable.StartState())
	default:
		return fmt.Errorf("sleep.Run: next step is not sfn chainable: got %T", step.Next)
	}

	return err
}

func (s *Lambda) SetOpts(opts *LambdaOpts) {
	s.LambdaOpts = opts
}
