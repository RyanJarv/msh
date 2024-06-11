package awsapicallviacloudtrail

type ResponseElements struct {
    HealthCheck HealthCheck `json:"healthCheck,omitempty"`
    Instances []ResponseElementsItem `json:"instances,omitempty"`
    DNSName string `json:"dNSName,omitempty"`
}

func (r *ResponseElements) SetHealthCheck(healthCheck HealthCheck) {
    r.HealthCheck = healthCheck
}

func (r *ResponseElements) SetInstances(instances []ResponseElementsItem) {
    r.Instances = instances
}

func (r *ResponseElements) SetDNSName(dNSName string) {
    r.DNSName = dNSName
}
