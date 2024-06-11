package awsapicallviacloudtrail

type ResponseElements struct {
    RetentionPeriod RetentionPeriod_1 `json:"retentionPeriod,omitempty"`
    ChannelArn string `json:"channelArn,omitempty"`
    ChannelName string `json:"channelName,omitempty"`
    DatasetArn string `json:"datasetArn,omitempty"`
    DatasetName string `json:"datasetName,omitempty"`
    DatastoreArn string `json:"datastoreArn,omitempty"`
    DatastoreIndexArn string `json:"datastoreIndexArn,omitempty"`
    DatastoreIndexName string `json:"datastoreIndexName,omitempty"`
    DatastoreName string `json:"datastoreName,omitempty"`
    PipelineArn string `json:"pipelineArn,omitempty"`
    PipelineName string `json:"pipelineName,omitempty"`
    ReprocessingId string `json:"reprocessingId,omitempty"`
    VersionId string `json:"versionId,omitempty"`
}

func (r *ResponseElements) SetRetentionPeriod(retentionPeriod RetentionPeriod_1) {
    r.RetentionPeriod = retentionPeriod
}

func (r *ResponseElements) SetChannelArn(channelArn string) {
    r.ChannelArn = channelArn
}

func (r *ResponseElements) SetChannelName(channelName string) {
    r.ChannelName = channelName
}

func (r *ResponseElements) SetDatasetArn(datasetArn string) {
    r.DatasetArn = datasetArn
}

func (r *ResponseElements) SetDatasetName(datasetName string) {
    r.DatasetName = datasetName
}

func (r *ResponseElements) SetDatastoreArn(datastoreArn string) {
    r.DatastoreArn = datastoreArn
}

func (r *ResponseElements) SetDatastoreIndexArn(datastoreIndexArn string) {
    r.DatastoreIndexArn = datastoreIndexArn
}

func (r *ResponseElements) SetDatastoreIndexName(datastoreIndexName string) {
    r.DatastoreIndexName = datastoreIndexName
}

func (r *ResponseElements) SetDatastoreName(datastoreName string) {
    r.DatastoreName = datastoreName
}

func (r *ResponseElements) SetPipelineArn(pipelineArn string) {
    r.PipelineArn = pipelineArn
}

func (r *ResponseElements) SetPipelineName(pipelineName string) {
    r.PipelineName = pipelineName
}

func (r *ResponseElements) SetReprocessingId(reprocessingId string) {
    r.ReprocessingId = reprocessingId
}

func (r *ResponseElements) SetVersionId(versionId string) {
    r.VersionId = versionId
}
