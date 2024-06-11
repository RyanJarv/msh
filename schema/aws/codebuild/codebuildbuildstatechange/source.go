package codebuildbuildstatechange

type Source struct {
    Auth Auth `json:"auth,omitempty"`
    Buildspec string `json:"buildspec,omitempty"`
    Location string `json:"location,omitempty"`
    SourceType string `json:"type"`
}

func (s *Source) SetAuth(auth Auth) {
    s.Auth = auth
}

func (s *Source) SetBuildspec(buildspec string) {
    s.Buildspec = buildspec
}

func (s *Source) SetLocation(location string) {
    s.Location = location
}

func (s *Source) SetSourceType(sourceType string) {
    s.SourceType = sourceType
}
