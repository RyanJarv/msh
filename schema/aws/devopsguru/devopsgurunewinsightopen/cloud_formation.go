package devopsgurunewinsightopen

type CloudFormation struct {
    StackNames []string `json:"stackNames,omitempty"`
}

func (c *CloudFormation) SetStackNames(stackNames []string) {
    c.StackNames = stackNames
}
