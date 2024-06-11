package voiceidupdatesessionaction

type SystemAttributes struct {
    AwsConnectOriginalContactArn string `json:"aws-connect-OriginalContactArn,omitempty"`
}

func (s *SystemAttributes) SetAwsConnectOriginalContactArn(awsConnectOriginalContactArn string) {
    s.AwsConnectOriginalContactArn = awsConnectOriginalContactArn
}
