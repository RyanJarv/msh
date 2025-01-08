package utils

import (
	"fmt"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/types"
)

func EnsureChain(scope constructs.Construct, next types.CdkStep, f func(next sfn.IChainable) (sfn.IChainable, error)) (sfn.IChainable, error) {
	if next == nil {
		return sfn.Chain_Start(sfn.NewSucceed(scope, jsii.String("sfn-succeed"), &sfn.SucceedProps{})), nil
	} else if v, ok := next.(sfn.IChainable); ok {
		return f(v)
	} else {
		return nil, fmt.Errorf("sfn.Run: next step is not sfn chainable: got %T", next)
	}
}
