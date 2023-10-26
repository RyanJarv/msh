package utils

import sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"

func OptionalChain(current sfn.INextable, next interface{}) sfn.INextable {
	if next != nil {
		chain, ok := next.(sfn.IChainable)
		if !ok {
			return current
		}

		current = current.Next(chain)
	}
	return current
}
