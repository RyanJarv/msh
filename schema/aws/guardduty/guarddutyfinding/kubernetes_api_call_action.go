package guarddutyfinding

type KubernetesApiCallAction struct {
    RemoteIpDetails RemoteIpDetails_2 `json:"remoteIpDetails,omitempty"`
    Parameters string `json:"parameters,omitempty"`
    RequestUri string `json:"requestUri,omitempty"`
    SourceIPs []string `json:"sourceIPs,omitempty"`
    StatusCode float64 `json:"statusCode,omitempty"`
    UserAgent string `json:"userAgent,omitempty"`
    Verb string `json:"verb,omitempty"`
}

func (k *KubernetesApiCallAction) SetRemoteIpDetails(remoteIpDetails RemoteIpDetails_2) {
    k.RemoteIpDetails = remoteIpDetails
}

func (k *KubernetesApiCallAction) SetParameters(parameters string) {
    k.Parameters = parameters
}

func (k *KubernetesApiCallAction) SetRequestUri(requestUri string) {
    k.RequestUri = requestUri
}

func (k *KubernetesApiCallAction) SetSourceIPs(sourceIPs []string) {
    k.SourceIPs = sourceIPs
}

func (k *KubernetesApiCallAction) SetStatusCode(statusCode float64) {
    k.StatusCode = statusCode
}

func (k *KubernetesApiCallAction) SetUserAgent(userAgent string) {
    k.UserAgent = userAgent
}

func (k *KubernetesApiCallAction) SetVerb(verb string) {
    k.Verb = verb
}
