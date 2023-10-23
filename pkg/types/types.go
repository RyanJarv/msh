package types

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"io"
)

type IStep interface {
	Name() string
}

type CdkStep interface {
	CdkStep(stack awscdk.Stack)
}

type SfnStep interface {
	SfnHook(stack awscdk.Stack, chain sfn.Chain) sfn.Chain
}

type Process interface {
	io.ReaderFrom
	io.WriterTo
	Run() error
}

type Deployable interface {
	Deploy() error
}
