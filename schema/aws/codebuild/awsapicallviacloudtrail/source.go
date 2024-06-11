package awsapicallviacloudtrail

type Source struct {
    Auth Auth `json:"auth,omitempty"`
    Buildspec string `json:"buildspec,omitempty"`
    InsecureSsl bool `json:"insecureSsl,omitempty"`
    Location string `json:"location,omitempty"`
    ReportBuildStatus bool `json:"reportBuildStatus,omitempty"`
    SourceType string `json:"type,omitempty"`
}

func (s *Source) SetAuth(auth Auth) {
    s.Auth = auth
}

func (s *Source) SetBuildspec(buildspec string) {
    s.Buildspec = buildspec
}

func (s *Source) SetInsecureSsl(insecureSsl bool) {
    s.InsecureSsl = insecureSsl
}

func (s *Source) SetLocation(location string) {
    s.Location = location
}

func (s *Source) SetReportBuildStatus(reportBuildStatus bool) {
    s.ReportBuildStatus = reportBuildStatus
}

func (s *Source) SetSourceType(sourceType string) {
    s.SourceType = sourceType
}
