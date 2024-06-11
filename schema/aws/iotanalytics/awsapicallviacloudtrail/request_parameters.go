package awsapicallviacloudtrail

type RequestParameters struct {
    ChannelStorage ChannelStorage `json:"channelStorage,omitempty"`
    DatastoreStorage ChannelStorage `json:"datastoreStorage,omitempty"`
    LoggingOptions LoggingOptions `json:"loggingOptions,omitempty"`
    PipelineActivity PipelineActivity `json:"pipelineActivity,omitempty"`
    QueryAction QueryAction `json:"queryAction,omitempty"`
    RetentionPeriod RetentionPeriod `json:"retentionPeriod,omitempty"`
    Settings Settings `json:"settings,omitempty"`
    ChannelName string `json:"channelName,omitempty"`
    DatasetName string `json:"datasetName,omitempty"`
    DatastoreIndexName string `json:"datastoreIndexName,omitempty"`
    DatastoreName string `json:"datastoreName,omitempty"`
    EndTime string `json:"endTime,omitempty"`
    MaxMessages float64 `json:"maxMessages,omitempty"`
    Payloads []RequestParametersItem `json:"payloads,omitempty"`
    PipelineName string `json:"pipelineName,omitempty"`
    ReprocessingId string `json:"reprocessingId,omitempty"`
    ResourceArn string `json:"resourceArn,omitempty"`
    StartTime string `json:"startTime,omitempty"`
    TagKeys []string `json:"tagKeys,omitempty"`
    Tags []RequestParametersItem_1 `json:"tags,omitempty"`
    VersionId string `json:"versionId,omitempty"`
}

func (r *RequestParameters) SetChannelStorage(channelStorage ChannelStorage) {
    r.ChannelStorage = channelStorage
}

func (r *RequestParameters) SetDatastoreStorage(datastoreStorage ChannelStorage) {
    r.DatastoreStorage = datastoreStorage
}

func (r *RequestParameters) SetLoggingOptions(loggingOptions LoggingOptions) {
    r.LoggingOptions = loggingOptions
}

func (r *RequestParameters) SetPipelineActivity(pipelineActivity PipelineActivity) {
    r.PipelineActivity = pipelineActivity
}

func (r *RequestParameters) SetQueryAction(queryAction QueryAction) {
    r.QueryAction = queryAction
}

func (r *RequestParameters) SetRetentionPeriod(retentionPeriod RetentionPeriod) {
    r.RetentionPeriod = retentionPeriod
}

func (r *RequestParameters) SetSettings(settings Settings) {
    r.Settings = settings
}

func (r *RequestParameters) SetChannelName(channelName string) {
    r.ChannelName = channelName
}

func (r *RequestParameters) SetDatasetName(datasetName string) {
    r.DatasetName = datasetName
}

func (r *RequestParameters) SetDatastoreIndexName(datastoreIndexName string) {
    r.DatastoreIndexName = datastoreIndexName
}

func (r *RequestParameters) SetDatastoreName(datastoreName string) {
    r.DatastoreName = datastoreName
}

func (r *RequestParameters) SetEndTime(endTime string) {
    r.EndTime = endTime
}

func (r *RequestParameters) SetMaxMessages(maxMessages float64) {
    r.MaxMessages = maxMessages
}

func (r *RequestParameters) SetPayloads(payloads []RequestParametersItem) {
    r.Payloads = payloads
}

func (r *RequestParameters) SetPipelineName(pipelineName string) {
    r.PipelineName = pipelineName
}

func (r *RequestParameters) SetReprocessingId(reprocessingId string) {
    r.ReprocessingId = reprocessingId
}

func (r *RequestParameters) SetResourceArn(resourceArn string) {
    r.ResourceArn = resourceArn
}

func (r *RequestParameters) SetStartTime(startTime string) {
    r.StartTime = startTime
}

func (r *RequestParameters) SetTagKeys(tagKeys []string) {
    r.TagKeys = tagKeys
}

func (r *RequestParameters) SetTags(tags []RequestParametersItem_1) {
    r.Tags = tags
}

func (r *RequestParameters) SetVersionId(versionId string) {
    r.VersionId = versionId
}
