package sagemakeralgorithmstatechange

type SageMakerAlgorithmStateChange struct {
    AlgorithmArn string `json:"AlgorithmArn,omitempty"`
    AlgorithmDescription string `json:"AlgorithmDescription,omitempty"`
    AlgorithmName string `json:"AlgorithmName,omitempty"`
    CertifyForMarketplace bool `json:"CertifyForMarketplace,omitempty"`
    Tags []Tags `json:"Tags,omitempty"`
}

func (s *SageMakerAlgorithmStateChange) SetAlgorithmArn(algorithmArn string) {
    s.AlgorithmArn = algorithmArn
}

func (s *SageMakerAlgorithmStateChange) SetAlgorithmDescription(algorithmDescription string) {
    s.AlgorithmDescription = algorithmDescription
}

func (s *SageMakerAlgorithmStateChange) SetAlgorithmName(algorithmName string) {
    s.AlgorithmName = algorithmName
}

func (s *SageMakerAlgorithmStateChange) SetCertifyForMarketplace(certifyForMarketplace bool) {
    s.CertifyForMarketplace = certifyForMarketplace
}

func (s *SageMakerAlgorithmStateChange) SetTags(tags []Tags) {
    s.Tags = tags
}
