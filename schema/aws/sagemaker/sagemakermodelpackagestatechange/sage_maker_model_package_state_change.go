package sagemakermodelpackagestatechange

type SageMakerModelPackageStateChange struct {
    Tags []Tags `json:"Tags,omitempty"`
    CertifyForMarketplace bool `json:"CertifyForMarketplace,omitempty"`
    ModelPackageArn string `json:"ModelPackageArn,omitempty"`
    ModelPackageDescription string `json:"ModelPackageDescription,omitempty"`
    ModelPackageName string `json:"ModelPackageName,omitempty"`
}

func (s *SageMakerModelPackageStateChange) SetTags(tags []Tags) {
    s.Tags = tags
}

func (s *SageMakerModelPackageStateChange) SetCertifyForMarketplace(certifyForMarketplace bool) {
    s.CertifyForMarketplace = certifyForMarketplace
}

func (s *SageMakerModelPackageStateChange) SetModelPackageArn(modelPackageArn string) {
    s.ModelPackageArn = modelPackageArn
}

func (s *SageMakerModelPackageStateChange) SetModelPackageDescription(modelPackageDescription string) {
    s.ModelPackageDescription = modelPackageDescription
}

func (s *SageMakerModelPackageStateChange) SetModelPackageName(modelPackageName string) {
    s.ModelPackageName = modelPackageName
}
