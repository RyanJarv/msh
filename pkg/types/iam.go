package types

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/samber/lo"
)

type IamPolicyStatement struct {
	Effect    string
	Action    []string
	Resource  []string   `json:",omitempty"`
	Principal *Principal `json:",omitempty"`
}

type Principal struct {
	Service string `json:",omitempty"`
	AWS     string `json:",omitempty"`
}

type IamPolicy struct {
	Name      string `json:"-"`
	Version   string
	Statement []IamPolicyStatement
}

func (p *IamPolicy) ToString() *string {
	return aws.String(string(lo.Must(json.Marshal(p))))
}
