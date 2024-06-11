package awsapicallviacloudtrail

type Settings struct {
    ChannelIdentification bool `json:"channelIdentification,omitempty"`
    VocabularyName string `json:"vocabularyName,omitempty"`
}

func (s *Settings) SetChannelIdentification(channelIdentification bool) {
    s.ChannelIdentification = channelIdentification
}

func (s *Settings) SetVocabularyName(vocabularyName string) {
    s.VocabularyName = vocabularyName
}
