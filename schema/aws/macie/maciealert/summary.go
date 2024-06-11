package maciealert

type Summary struct {
    DLPRisk DLP_risk `json:"DLP-risk,omitempty"`
    Owner Owner `json:"Owner,omitempty"`
    IP IP `json:"IP,omitempty"`
    ACL ACL `json:"ACL,omitempty"`
    Themes Themes `json:"Themes,omitempty"`
    Timestamps Timestamps `json:"Timestamps,omitempty"`
    Events Events `json:"Events,omitempty"`
    Bucket Bucket `json:"Bucket,omitempty"`
    RecipientAccountId RecipientAccountId `json:"recipientAccountId,omitempty"`
    Object Object `json:"Object,omitempty"`
    Location Location `json:"Location,omitempty"`
    SourceARN string `json:"Source ARN,omitempty"`
    Description string `json:"Description,omitempty"`
    TimeRange []SummaryItem `json:"Time Range,omitempty"`
    RecordCount float64 `json:"Record-Count,omitempty"`
    EventCount float64 `json:"Event-Count,omitempty"`
}

func (s *Summary) SetDLPRisk(dLPRisk DLP_risk) {
    s.DLPRisk = dLPRisk
}

func (s *Summary) SetOwner(owner Owner) {
    s.Owner = owner
}

func (s *Summary) SetIP(iP IP) {
    s.IP = iP
}

func (s *Summary) SetACL(aCL ACL) {
    s.ACL = aCL
}

func (s *Summary) SetThemes(themes Themes) {
    s.Themes = themes
}

func (s *Summary) SetTimestamps(timestamps Timestamps) {
    s.Timestamps = timestamps
}

func (s *Summary) SetEvents(events Events) {
    s.Events = events
}

func (s *Summary) SetBucket(bucket Bucket) {
    s.Bucket = bucket
}

func (s *Summary) SetRecipientAccountId(recipientAccountId RecipientAccountId) {
    s.RecipientAccountId = recipientAccountId
}

func (s *Summary) SetObject(object Object) {
    s.Object = object
}

func (s *Summary) SetLocation(location Location) {
    s.Location = location
}

func (s *Summary) SetSourceARN(sourceARN string) {
    s.SourceARN = sourceARN
}

func (s *Summary) SetDescription(description string) {
    s.Description = description
}

func (s *Summary) SetTimeRange(timeRange []SummaryItem) {
    s.TimeRange = timeRange
}

func (s *Summary) SetRecordCount(recordCount float64) {
    s.RecordCount = recordCount
}

func (s *Summary) SetEventCount(eventCount float64) {
    s.EventCount = eventCount
}
