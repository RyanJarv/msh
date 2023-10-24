package types

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"io"
)

type IStep interface {
	Name() string
}

type CdkStep interface {
	//CdkStep(stack awscdk.Stack)
	Run(awscdk.Stack, interface{}) (interface{}, error)
}

type Process interface {
	io.ReaderFrom
	io.WriterTo
	Run() error
}

type Deployable interface {
	Deploy() error
}
