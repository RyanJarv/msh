package awsapicallviacloudtrail

type Source_1 struct {
    Auth Auth `json:"auth,omitempty"`
    GitSubmodulesConfig GitSubmodulesConfig `json:"gitSubmodulesConfig,omitempty"`
    Buildspec string `json:"buildspec,omitempty"`
    GitCloneDepth float64 `json:"gitCloneDepth,omitempty"`
    InsecureSsl bool `json:"insecureSsl,omitempty"`
    Location string `json:"location,omitempty"`
    ReportBuildStatus bool `json:"reportBuildStatus,omitempty"`
    Source_1Type string `json:"type,omitempty"`
}

func (s *Source_1) SetAuth(auth Auth) {
    s.Auth = auth
}

func (s *Source_1) SetGitSubmodulesConfig(gitSubmodulesConfig GitSubmodulesConfig) {
    s.GitSubmodulesConfig = gitSubmodulesConfig
}

func (s *Source_1) SetBuildspec(buildspec string) {
    s.Buildspec = buildspec
}

func (s *Source_1) SetGitCloneDepth(gitCloneDepth float64) {
    s.GitCloneDepth = gitCloneDepth
}

func (s *Source_1) SetInsecureSsl(insecureSsl bool) {
    s.InsecureSsl = insecureSsl
}

func (s *Source_1) SetLocation(location string) {
    s.Location = location
}

func (s *Source_1) SetReportBuildStatus(reportBuildStatus bool) {
    s.ReportBuildStatus = reportBuildStatus
}

func (s *Source_1) SetSource1Type(source1Type string) {
    s.Source_1Type = source1Type
}
