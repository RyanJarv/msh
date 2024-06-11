package codebuildbuildphasechange

type Artifact struct {
    Location string `json:"location"`
    Md5sum string `json:"md5sum,omitempty"`
    Sha256sum string `json:"sha256sum,omitempty"`
}

func (a *Artifact) SetLocation(location string) {
    a.Location = location
}

func (a *Artifact) SetMd5sum(md5sum string) {
    a.Md5sum = md5sum
}

func (a *Artifact) SetSha256sum(sha256sum string) {
    a.Sha256sum = sha256sum
}
