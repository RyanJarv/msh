package lambda

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/jsii-runtime-go"
	"log"
	"os"
)

func NewLambda(args []string) (*LambdaCmd, error) {
	if len(args) < 1 {
		log.Fatalf("usage: %s <script>", os.Args[0])
	}

	script, err := os.ReadFile(args[0])
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

func (s *LambdaCmd) Name() string { return "lambda" }

func (s *LambdaCmd) Run(stack awscdk.Stack, last interface{}) (interface{}, error) {
	chain, ok := last.(awsstepfunctions.INextable)
	if !ok {
		return nil, fmt.Errorf("last step must be statemachine chain")
	}

	s.function = awslambda.NewFunction(stack, jsii.String(flag.Arg(0)), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PYTHON_3_11(),
		Handler: jsii.String("index.lambda_handler"),
		Code:    awslambda.Code_FromInline(jsii.String(s.Script)),
	})
	if chain == nil {
		return chain, fmt.Errorf("end of chain not found")
	}

	lambda := tasks.NewLambdaInvoke(stack, jsii.String(fmt.Sprintf("invoke %s %d", flag.Arg(0), os.Getpid())), &tasks.LambdaInvokeProps{
		LambdaFunction: s.function,
		OutputPath:     jsii.String("$.Payload"),
	})

	return chain.Next(lambda), nil
}
