package objectdeleted

type Object struct {
    Etag string `json:"etag"`
    Key string `json:"key"`
    Sequencer string `json:"sequencer"`
    VersionId string `json:"version-id,omitempty"`
}

func (o *Object) SetEtag(etag string) {
    o.Etag = etag
}

func (o *Object) SetKey(key string) {
    o.Key = key
}

func (o *Object) SetSequencer(sequencer string) {
    o.Sequencer = sequencer
}

func (o *Object) SetVersionId(versionId string) {
    o.VersionId = versionId
}
