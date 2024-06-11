package awshealthevent

type EventDescription struct {
    LatestDescription string `json:"latestDescription,omitempty"`
    Language string `json:"language,omitempty"`
}

func (e *EventDescription) SetLatestDescription(latestDescription string) {
    e.LatestDescription = latestDescription
}

func (e *EventDescription) SetLanguage(language string) {
    e.Language = language
}
