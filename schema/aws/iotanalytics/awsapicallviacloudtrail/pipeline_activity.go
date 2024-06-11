package awsapicallviacloudtrail

type PipelineActivity struct {
    Math Math `json:"math,omitempty"`
}

func (p *PipelineActivity) SetMath(math Math) {
    p.Math = math
}
