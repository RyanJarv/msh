package awsapicallviacloudtrail

type RequestParametersItem struct {
    Expression string `json:"expression"`
    RequestParametersItemType string `json:"type"`
}

func (r *RequestParametersItem) SetExpression(expression string) {
    r.Expression = expression
}

func (r *RequestParametersItem) SetRequestParametersItemType(requestParametersItemType string) {
    r.RequestParametersItemType = requestParametersItemType
}
