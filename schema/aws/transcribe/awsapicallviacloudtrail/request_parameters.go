package awsapicallviacloudtrail

type RequestParameters struct {
    Media Media `json:"media,omitempty"`
    Settings Settings `json:"settings,omitempty"`
    LanguageCode string `json:"languageCode,omitempty"`
    MediaFormat string `json:"mediaFormat,omitempty"`
    MediaSampleRateHertz float64 `json:"mediaSampleRateHertz,omitempty"`
    TranscriptionJobName string `json:"transcriptionJobName,omitempty"`
    VocabularyName string `json:"vocabularyName,omitempty"`
}

func (r *RequestParameters) SetMedia(media Media) {
    r.Media = media
}

func (r *RequestParameters) SetSettings(settings Settings) {
    r.Settings = settings
}

func (r *RequestParameters) SetLanguageCode(languageCode string) {
    r.LanguageCode = languageCode
}

func (r *RequestParameters) SetMediaFormat(mediaFormat string) {
    r.MediaFormat = mediaFormat
}

func (r *RequestParameters) SetMediaSampleRateHertz(mediaSampleRateHertz float64) {
    r.MediaSampleRateHertz = mediaSampleRateHertz
}

func (r *RequestParameters) SetTranscriptionJobName(transcriptionJobName string) {
    r.TranscriptionJobName = transcriptionJobName
}

func (r *RequestParameters) SetVocabularyName(vocabularyName string) {
    r.VocabularyName = vocabularyName
}
