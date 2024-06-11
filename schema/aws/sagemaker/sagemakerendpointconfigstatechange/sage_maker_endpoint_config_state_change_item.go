package sagemakerendpointconfigstatechange

type SageMakerEndpointConfigStateChangeItem struct {
    ModelName string `json:"ModelName"`
    VariantName string `json:"VariantName"`
    InitialInstanceCount float64 `json:"InitialInstanceCount"`
    InstanceType string `json:"InstanceType"`
    InitialVariantWeight float64 `json:"InitialVariantWeight"`
}

func (s *SageMakerEndpointConfigStateChangeItem) SetModelName(modelName string) {
    s.ModelName = modelName
}

func (s *SageMakerEndpointConfigStateChangeItem) SetVariantName(variantName string) {
    s.VariantName = variantName
}

func (s *SageMakerEndpointConfigStateChangeItem) SetInitialInstanceCount(initialInstanceCount float64) {
    s.InitialInstanceCount = initialInstanceCount
}

func (s *SageMakerEndpointConfigStateChangeItem) SetInstanceType(instanceType string) {
    s.InstanceType = instanceType
}

func (s *SageMakerEndpointConfigStateChangeItem) SetInitialVariantWeight(initialVariantWeight float64) {
    s.InitialVariantWeight = initialVariantWeight
}
