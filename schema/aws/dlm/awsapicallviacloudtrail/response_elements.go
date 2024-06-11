package awsapicallviacloudtrail

type ResponseElements struct {
    PolicyId string `json:"PolicyId"`
}

func (r *ResponseElements) SetPolicyId(policyId string) {
    r.PolicyId = policyId
}
