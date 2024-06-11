package awsxrayinsightupdate

type AWSXRayInsightUpdateItem struct {
    ServiceId ServiceId `json:"ServiceId"`
}

func (a *AWSXRayInsightUpdateItem) SetServiceId(serviceId ServiceId) {
    a.ServiceId = serviceId
}
