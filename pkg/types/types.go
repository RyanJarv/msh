package types

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"io"
)

type Registry map[string]IStep

type IStep interface {
	Name() string
}

type CdkStep interface {
	Compile(awscdk.Stack, interface{}) ([]interface{}, error)
}

type CdkStepFanOut interface {
	Compile(awscdk.Stack, []interface{}) ([]interface{}, error)
}

type Process interface {
	io.ReaderFrom
	io.WriterTo
	Run() error
}

type Deployable interface {
	Deploy() error
}
