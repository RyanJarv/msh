package awsapicallviacloudtrail

type TransformInput struct {
    DataSource DataSource `json:"dataSource,omitempty"`
    CompressionType string `json:"compressionType,omitempty"`
    ContentType string `json:"contentType,omitempty"`
}

func (t *TransformInput) SetDataSource(dataSource DataSource) {
    t.DataSource = dataSource
}

func (t *TransformInput) SetCompressionType(compressionType string) {
    t.CompressionType = compressionType
}

func (t *TransformInput) SetContentType(contentType string) {
    t.ContentType = contentType
}
