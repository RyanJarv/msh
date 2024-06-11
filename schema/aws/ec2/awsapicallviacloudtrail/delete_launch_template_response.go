package awsapicallviacloudtrail

type DeleteLaunchTemplateResponse struct {
    LaunchTemplate LaunchTemplate_1 `json:"launchTemplate,omitempty"`
    Xmlns string `json:"xmlns,omitempty"`
    RequestId string `json:"requestId,omitempty"`
}

func (d *DeleteLaunchTemplateResponse) SetLaunchTemplate(launchTemplate LaunchTemplate_1) {
    d.LaunchTemplate = launchTemplate
}

func (d *DeleteLaunchTemplateResponse) SetXmlns(xmlns string) {
    d.Xmlns = xmlns
}

func (d *DeleteLaunchTemplateResponse) SetRequestId(requestId string) {
    d.RequestId = requestId
}
