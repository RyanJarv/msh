package sagemakerendpointconfigstatechange

type SageMakerEndpointConfigStateChange struct {
    Tags []Tags `json:"Tags,omitempty"`
    ProductionVariants []SageMakerEndpointConfigStateChangeItem `json:"ProductionVariants"`
    EndpointConfigName string `json:"EndpointConfigName"`
    CreationTime float64 `json:"CreationTime"`
    EndpointConfigArn string `json:"EndpointConfigArn"`
}

func (s *SageMakerEndpointConfigStateChange) SetTags(tags []Tags) {
    s.Tags = tags
}

func (s *SageMakerEndpointConfigStateChange) SetProductionVariants(productionVariants []SageMakerEndpointConfigStateChangeItem) {
    s.ProductionVariants = productionVariants
}

func (s *SageMakerEndpointConfigStateChange) SetEndpointConfigName(endpointConfigName string) {
    s.EndpointConfigName = endpointConfigName
}

func (s *SageMakerEndpointConfigStateChange) SetCreationTime(creationTime float64) {
    s.CreationTime = creationTime
}

func (s *SageMakerEndpointConfigStateChange) SetEndpointConfigArn(endpointConfigArn string) {
    s.EndpointConfigArn = endpointConfigArn
}
