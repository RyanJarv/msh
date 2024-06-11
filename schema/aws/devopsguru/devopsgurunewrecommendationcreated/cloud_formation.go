package devopsgurunewrecommendationcreated

type CloudFormation struct {
    StackNames []string `json:"stackNames"`
}

func (c *CloudFormation) SetStackNames(stackNames []string) {
    c.StackNames = stackNames
}
