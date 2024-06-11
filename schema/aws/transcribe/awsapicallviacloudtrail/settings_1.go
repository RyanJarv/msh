package awsapicallviacloudtrail

type Settings_1 struct {
    ChannelIdentification bool `json:"channelIdentification,omitempty"`
    VocabularyName string `json:"vocabularyName,omitempty"`
}

func (s *Settings_1) SetChannelIdentification(channelIdentification bool) {
    s.ChannelIdentification = channelIdentification
}

func (s *Settings_1) SetVocabularyName(vocabularyName string) {
    s.VocabularyName = vocabularyName
}
