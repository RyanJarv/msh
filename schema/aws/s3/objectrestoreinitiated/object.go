package objectrestoreinitiated

type Object struct {
    Etag string `json:"etag"`
    Key string `json:"key"`
    Size float64 `json:"size"`
    VersionId string `json:"version-id,omitempty"`
}

func (o *Object) SetEtag(etag string) {
    o.Etag = etag
}

func (o *Object) SetKey(key string) {
    o.Key = key
}

func (o *Object) SetSize(size float64) {
    o.Size = size
}

func (o *Object) SetVersionId(versionId string) {
    o.VersionId = versionId
}
