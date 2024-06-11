package awsapicallviacloudtrail

type ResponseElements struct {
    EnvironmentId string `json:"EnvironmentId"`
    SessionId string `json:"SessionId,omitempty"`
    Status string `json:"Status,omitempty"`
    StreamUrl string `json:"StreamUrl,omitempty"`
    TokenValue string `json:"TokenValue,omitempty"`
}

func (r *ResponseElements) SetEnvironmentId(environmentId string) {
    r.EnvironmentId = environmentId
}

func (r *ResponseElements) SetSessionId(sessionId string) {
    r.SessionId = sessionId
}

func (r *ResponseElements) SetStatus(status string) {
    r.Status = status
}

func (r *ResponseElements) SetStreamUrl(streamUrl string) {
    r.StreamUrl = streamUrl
}

func (r *ResponseElements) SetTokenValue(tokenValue string) {
    r.TokenValue = tokenValue
}
