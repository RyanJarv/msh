package awsapicallviacloudtrail

type AdditionalEventData struct {
    Capabilities []string `json:"capabilities,omitempty"`
    Clone bool `json:"clone,omitempty"`
    DataTransferred bool `json:"dataTransferred,omitempty"`
    Protocol string `json:"protocol,omitempty"`
    RepositoryId string `json:"repositoryId,omitempty"`
    RepositoryName string `json:"repositoryName,omitempty"`
    Shallow bool `json:"shallow,omitempty"`
}

func (a *AdditionalEventData) SetCapabilities(capabilities []string) {
    a.Capabilities = capabilities
}

func (a *AdditionalEventData) SetClone(clone bool) {
    a.Clone = clone
}

func (a *AdditionalEventData) SetDataTransferred(dataTransferred bool) {
    a.DataTransferred = dataTransferred
}

func (a *AdditionalEventData) SetProtocol(protocol string) {
    a.Protocol = protocol
}

func (a *AdditionalEventData) SetRepositoryId(repositoryId string) {
    a.RepositoryId = repositoryId
}

func (a *AdditionalEventData) SetRepositoryName(repositoryName string) {
    a.RepositoryName = repositoryName
}

func (a *AdditionalEventData) SetShallow(shallow bool) {
    a.Shallow = shallow
}
