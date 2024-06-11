package sagemakertrainingjobstatechange

type SageMakerTrainingJobStateChangeItem struct {
    DataSource DataSource `json:"DataSource,omitempty"`
    ChannelName string `json:"ChannelName,omitempty"`
    ContentType string `json:"ContentType,omitempty"`
    RecordWrapperType string `json:"RecordWrapperType,omitempty"`
    CompressionType string `json:"CompressionType,omitempty"`
}

func (s *SageMakerTrainingJobStateChangeItem) SetDataSource(dataSource DataSource) {
    s.DataSource = dataSource
}

func (s *SageMakerTrainingJobStateChangeItem) SetChannelName(channelName string) {
    s.ChannelName = channelName
}

func (s *SageMakerTrainingJobStateChangeItem) SetContentType(contentType string) {
    s.ContentType = contentType
}

func (s *SageMakerTrainingJobStateChangeItem) SetRecordWrapperType(recordWrapperType string) {
    s.RecordWrapperType = recordWrapperType
}

func (s *SageMakerTrainingJobStateChangeItem) SetCompressionType(compressionType string) {
    s.CompressionType = compressionType
}
