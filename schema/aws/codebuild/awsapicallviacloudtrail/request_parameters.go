package awsapicallviacloudtrail

type RequestParameters struct {
    AwsActId string `json:"awsActId"`
    InvokedBy string `json:"invokedBy,omitempty"`
    PayerId string `json:"payerId"`
    UserArn string `json:"userArn"`
}

func (r *RequestParameters) SetAwsActId(awsActId string) {
    r.AwsActId = awsActId
}

func (r *RequestParameters) SetInvokedBy(invokedBy string) {
    r.InvokedBy = invokedBy
}

func (r *RequestParameters) SetPayerId(payerId string) {
    r.PayerId = payerId
}

func (r *RequestParameters) SetUserArn(userArn string) {
    r.UserArn = userArn
}
