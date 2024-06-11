package sagemakerendpointstatechange

type SageMakerEndpointStateChange struct {
    Tags []Tags `json:"Tags,omitempty"`
    CreationTime float64 `json:"CreationTime,omitempty"`
    EndpointArn string `json:"EndpointArn,omitempty"`
    EndpointConfigName string `json:"EndpointConfigName,omitempty"`
    EndpointName string `json:"EndpointName,omitempty"`
    EndpointStatus string `json:"EndpointStatus,omitempty"`
    LastModifiedTime float64 `json:"LastModifiedTime,omitempty"`
}

func (s *SageMakerEndpointStateChange) SetTags(tags []Tags) {
    s.Tags = tags
}

func (s *SageMakerEndpointStateChange) SetCreationTime(creationTime float64) {
    s.CreationTime = creationTime
}

func (s *SageMakerEndpointStateChange) SetEndpointArn(endpointArn string) {
    s.EndpointArn = endpointArn
}

func (s *SageMakerEndpointStateChange) SetEndpointConfigName(endpointConfigName string) {
    s.EndpointConfigName = endpointConfigName
}

func (s *SageMakerEndpointStateChange) SetEndpointName(endpointName string) {
    s.EndpointName = endpointName
}

func (s *SageMakerEndpointStateChange) SetEndpointStatus(endpointStatus string) {
    s.EndpointStatus = endpointStatus
}

func (s *SageMakerEndpointStateChange) SetLastModifiedTime(lastModifiedTime float64) {
    s.LastModifiedTime = lastModifiedTime
}
