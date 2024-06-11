package awsapicallviacloudtrail

type WebIdFederationData struct {
    FederatedProvider string `json:"federatedProvider,omitempty"`
    Attributes map[string]string `json:"attributes,omitempty"`
}

func (w *WebIdFederationData) SetFederatedProvider(federatedProvider string) {
    w.FederatedProvider = federatedProvider
}

func (w *WebIdFederationData) SetAttributes(attributes map[string]string) {
    w.Attributes = attributes
}
