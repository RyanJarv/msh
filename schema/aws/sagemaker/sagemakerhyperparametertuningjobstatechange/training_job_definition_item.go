package sagemakerhyperparametertuningjobstatechange

type TrainingJobDefinitionItem struct {
    DataSource DataSource `json:"DataSource"`
    ChannelName string `json:"ChannelName"`
    ContentType string `json:"ContentType"`
    RecordWrapperType string `json:"RecordWrapperType"`
    CompressionType string `json:"CompressionType"`
}

func (t *TrainingJobDefinitionItem) SetDataSource(dataSource DataSource) {
    t.DataSource = dataSource
}

func (t *TrainingJobDefinitionItem) SetChannelName(channelName string) {
    t.ChannelName = channelName
}

func (t *TrainingJobDefinitionItem) SetContentType(contentType string) {
    t.ContentType = contentType
}

func (t *TrainingJobDefinitionItem) SetRecordWrapperType(recordWrapperType string) {
    t.RecordWrapperType = recordWrapperType
}

func (t *TrainingJobDefinitionItem) SetCompressionType(compressionType string) {
    t.CompressionType = compressionType
}
