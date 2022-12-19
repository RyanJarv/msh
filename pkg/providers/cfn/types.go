package cfn

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type Aws struct {
	Config aws.Config `json:"-"`
}

type Global struct {
	Aws     Aws
	Args    []string
	Project string
}

type RunTemplate interface {
	Parsable
	Runnable
}

type Parsable interface {
	Parse(map[string]string)
}

type Runnable interface {
	Run() error
}

type Parser struct {
	Content string

	Args []string

	context.Context `json:"-"`
	Global
	Runnable `json:"-"`
}
