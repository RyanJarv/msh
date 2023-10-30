package lambda

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"log"
	"os"
)

func New(args []string) (*LambdaCmd, error) {
	if len(args) > 2 {
		log.Fatalf("usage: %s <script>", os.Args[0])
	}
	flagset := flag.NewFlagSet("lambda", flag.ExitOnError)
	err := flagset.Parse(args)
	if err != nil {
		return nil, fmt.Errorf("parsing flags: %w", err)
	}

	path := flagset.Arg(flagset.NArg() - 1)

	script, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	if !bytes.Contains(script, []byte("lambda_handler")) {
		return nil, fmt.Errorf("script must contain a `lambda_handler` function")
	}

	return &LambdaCmd{
		Script: string(script),
		Args:   args,
	}, nil
}

type LambdaCmd struct {
	Script   string
	Args     []string
	function awslambda.Function
}

func (s LambdaCmd) GetName() string { return "lambda" }

func (s LambdaCmd) Compile(stack constructs.Construct, next interface{}) ([]interface{}, error) {
	s.function = awslambda.NewFunction(stack, jsii.String(s.GetName()), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PYTHON_3_11(),
		Handler: jsii.String("index.lambda_handler"),
		Code:    awslambda.Code_FromInline(jsii.String(s.Script)),
	})

	var this awsstepfunctions.INextable

	this = tasks.NewLambdaInvoke(stack, jsii.String(fmt.Sprintf("%s-invoke", s.GetName())), &tasks.LambdaInvokeProps{
		LambdaFunction: s.function,
		OutputPath:     jsii.String("$.Payload"),
	})

	if next != nil {
		chain, ok := next.(awsstepfunctions.IChainable)
		if !ok {
			return nil, fmt.Errorf("next step must be statemachine chain")
		}

		this = this.Next(chain)
	}

	return []interface{}{this}, nil
}
