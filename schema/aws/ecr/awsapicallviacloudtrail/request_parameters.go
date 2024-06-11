package awsapicallviacloudtrail

type RequestParameters struct {
    AcceptedMediaTypes []string `json:"acceptedMediaTypes"`
    RegistryId string `json:"registryId"`
    RepositoryName string `json:"repositoryName"`
    ImageIds []RequestParametersItem `json:"imageIds"`
}

func (r *RequestParameters) SetAcceptedMediaTypes(acceptedMediaTypes []string) {
    r.AcceptedMediaTypes = acceptedMediaTypes
}

func (r *RequestParameters) SetRegistryId(registryId string) {
    r.RegistryId = registryId
}

func (r *RequestParameters) SetRepositoryName(repositoryName string) {
    r.RepositoryName = repositoryName
}

func (r *RequestParameters) SetImageIds(imageIds []RequestParametersItem) {
    r.ImageIds = imageIds
}
