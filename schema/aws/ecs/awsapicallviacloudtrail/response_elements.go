package awsapicallviacloudtrail

type ResponseElements struct {
    Endpoint string `json:"endpoint,omitempty"`
    Failures interface{} `json:"failures,omitempty"`
    TelemetryEndpoint string `json:"telemetryEndpoint,omitempty"`
    Tasks []ResponseElementsItem `json:"tasks,omitempty"`
    Acknowledgment string `json:"acknowledgment,omitempty"`
}

func (r *ResponseElements) SetEndpoint(endpoint string) {
    r.Endpoint = endpoint
}

func (r *ResponseElements) SetFailures(failures interface{}) {
    r.Failures = failures
}

func (r *ResponseElements) SetTelemetryEndpoint(telemetryEndpoint string) {
    r.TelemetryEndpoint = telemetryEndpoint
}

func (r *ResponseElements) SetTasks(tasks []ResponseElementsItem) {
    r.Tasks = tasks
}

func (r *ResponseElements) SetAcknowledgment(acknowledgment string) {
    r.Acknowledgment = acknowledgment
}
