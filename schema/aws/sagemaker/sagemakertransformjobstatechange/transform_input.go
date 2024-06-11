package sagemakertransformjobstatechange

type TransformInput struct {
    DataSource DataSource `json:"DataSource"`
    ContentType string `json:"ContentType"`
    SplitType string `json:"SplitType"`
    CompressionType string `json:"CompressionType"`
}

func (t *TransformInput) SetDataSource(dataSource DataSource) {
    t.DataSource = dataSource
}

func (t *TransformInput) SetContentType(contentType string) {
    t.ContentType = contentType
}

func (t *TransformInput) SetSplitType(splitType string) {
    t.SplitType = splitType
}

func (t *TransformInput) SetCompressionType(compressionType string) {
    t.CompressionType = compressionType
}
