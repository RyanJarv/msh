package awsapicallviacloudtrail

type PrimaryContainer struct {
    Image string `json:"image,omitempty"`
    ModelDataUrl string `json:"modelDataUrl,omitempty"`
    ContainerHostname string `json:"containerHostname,omitempty"`
}

func (p *PrimaryContainer) SetImage(image string) {
    p.Image = image
}

func (p *PrimaryContainer) SetModelDataUrl(modelDataUrl string) {
    p.ModelDataUrl = modelDataUrl
}

func (p *PrimaryContainer) SetContainerHostname(containerHostname string) {
    p.ContainerHostname = containerHostname
}
