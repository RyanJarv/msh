package awsapicallviacloudtrail

type Webhook struct {
    BranchFilter string `json:"branchFilter,omitempty"`
    LastModifiedSecret string `json:"lastModifiedSecret,omitempty"`
    PayloadUrl string `json:"payloadUrl,omitempty"`
    Url string `json:"url,omitempty"`
}

func (w *Webhook) SetBranchFilter(branchFilter string) {
    w.BranchFilter = branchFilter
}

func (w *Webhook) SetLastModifiedSecret(lastModifiedSecret string) {
    w.LastModifiedSecret = lastModifiedSecret
}

func (w *Webhook) SetPayloadUrl(payloadUrl string) {
    w.PayloadUrl = payloadUrl
}

func (w *Webhook) SetUrl(url string) {
    w.Url = url
}
