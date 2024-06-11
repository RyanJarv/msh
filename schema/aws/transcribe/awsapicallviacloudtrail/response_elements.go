package awsapicallviacloudtrail

type ResponseElements struct {
    TranscriptionJob TranscriptionJob `json:"transcriptionJob,omitempty"`
    LanguageCode string `json:"languageCode,omitempty"`
    VocabularyName string `json:"vocabularyName,omitempty"`
    VocabularyState string `json:"vocabularyState,omitempty"`
}

func (r *ResponseElements) SetTranscriptionJob(transcriptionJob TranscriptionJob) {
    r.TranscriptionJob = transcriptionJob
}

func (r *ResponseElements) SetLanguageCode(languageCode string) {
    r.LanguageCode = languageCode
}

func (r *ResponseElements) SetVocabularyName(vocabularyName string) {
    r.VocabularyName = vocabularyName
}

func (r *ResponseElements) SetVocabularyState(vocabularyState string) {
    r.VocabularyState = vocabularyState
}
