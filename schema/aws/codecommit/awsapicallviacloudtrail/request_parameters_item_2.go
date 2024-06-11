package awsapicallviacloudtrail

type RequestParametersItem_2 struct {
    DestinationReference string `json:"destinationReference"`
    RepositoryName string `json:"repositoryName"`
    SourceReference string `json:"sourceReference"`
}

func (r *RequestParametersItem_2) SetDestinationReference(destinationReference string) {
    r.DestinationReference = destinationReference
}

func (r *RequestParametersItem_2) SetRepositoryName(repositoryName string) {
    r.RepositoryName = repositoryName
}

func (r *RequestParametersItem_2) SetSourceReference(sourceReference string) {
    r.SourceReference = sourceReference
}
