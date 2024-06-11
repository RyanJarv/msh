package recordingstatechange

type IVSRecordingStateChange struct {
    ChannelName string `json:"channel_name"`
    StreamId string `json:"stream_id"`
    RecordingStatus string `json:"recording_status"`
    RecordingStatusReason string `json:"recording_status_reason"`
    RecordingS3BucketName string `json:"recording_s3_bucket_name"`
    RecordingS3KeyPrefix string `json:"recording_s3_key_prefix"`
}

func (i *IVSRecordingStateChange) SetChannelName(channelName string) {
    i.ChannelName = channelName
}

func (i *IVSRecordingStateChange) SetStreamId(streamId string) {
    i.StreamId = streamId
}

func (i *IVSRecordingStateChange) SetRecordingStatus(recordingStatus string) {
    i.RecordingStatus = recordingStatus
}

func (i *IVSRecordingStateChange) SetRecordingStatusReason(recordingStatusReason string) {
    i.RecordingStatusReason = recordingStatusReason
}

func (i *IVSRecordingStateChange) SetRecordingS3BucketName(recordingS3BucketName string) {
    i.RecordingS3BucketName = recordingS3BucketName
}

func (i *IVSRecordingStateChange) SetRecordingS3KeyPrefix(recordingS3KeyPrefix string) {
    i.RecordingS3KeyPrefix = recordingS3KeyPrefix
}
