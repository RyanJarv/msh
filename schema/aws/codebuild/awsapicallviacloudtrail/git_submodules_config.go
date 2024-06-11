package awsapicallviacloudtrail

type GitSubmodulesConfig struct {
    FetchSubmodules bool `json:"fetchSubmodules,omitempty"`
}

func (g *GitSubmodulesConfig) SetFetchSubmodules(fetchSubmodules bool) {
    g.FetchSubmodules = fetchSubmodules
}
