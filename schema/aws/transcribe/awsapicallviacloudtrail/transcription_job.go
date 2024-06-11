package awsapicallviacloudtrail

type TranscriptionJob struct {
    Media Media_1 `json:"media,omitempty"`
    Settings Settings_1 `json:"settings,omitempty"`
    CreationTime string `json:"creationTime,omitempty"`
    LanguageCode string `json:"languageCode,omitempty"`
    MediaFormat string `json:"mediaFormat,omitempty"`
    MediaSampleRateHertz float64 `json:"mediaSampleRateHertz,omitempty"`
    TranscriptionJobName string `json:"transcriptionJobName,omitempty"`
    TranscriptionJobStatus string `json:"transcriptionJobStatus,omitempty"`
}

func (t *TranscriptionJob) SetMedia(media Media_1) {
    t.Media = media
}

func (t *TranscriptionJob) SetSettings(settings Settings_1) {
    t.Settings = settings
}

func (t *TranscriptionJob) SetCreationTime(creationTime string) {
    t.CreationTime = creationTime
}

func (t *TranscriptionJob) SetLanguageCode(languageCode string) {
    t.LanguageCode = languageCode
}

func (t *TranscriptionJob) SetMediaFormat(mediaFormat string) {
    t.MediaFormat = mediaFormat
}

func (t *TranscriptionJob) SetMediaSampleRateHertz(mediaSampleRateHertz float64) {
    t.MediaSampleRateHertz = mediaSampleRateHertz
}

func (t *TranscriptionJob) SetTranscriptionJobName(transcriptionJobName string) {
    t.TranscriptionJobName = transcriptionJobName
}

func (t *TranscriptionJob) SetTranscriptionJobStatus(transcriptionJobStatus string) {
    t.TranscriptionJobStatus = transcriptionJobStatus
}
