package transcribejobstatechange

type TranscribeJobStateChange struct {
    TranscriptionJobName string `json:"TranscriptionJobName"`
    TranscriptionJobStatus string `json:"TranscriptionJobStatus"`
}

func (t *TranscribeJobStateChange) SetTranscriptionJobName(transcriptionJobName string) {
    t.TranscriptionJobName = transcriptionJobName
}

func (t *TranscribeJobStateChange) SetTranscriptionJobStatus(transcriptionJobStatus string) {
    t.TranscriptionJobStatus = transcriptionJobStatus
}
