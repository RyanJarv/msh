package sagemakermodelstatechange

type PrimaryContainer struct {
    ContainerHostname string `json:"ContainerHostname"`
    ModelDataUrl string `json:"ModelDataUrl"`
    Image string `json:"Image"`
}

func (p *PrimaryContainer) SetContainerHostname(containerHostname string) {
    p.ContainerHostname = containerHostname
}

func (p *PrimaryContainer) SetModelDataUrl(modelDataUrl string) {
    p.ModelDataUrl = modelDataUrl
}

func (p *PrimaryContainer) SetImage(image string) {
    p.Image = image
}
