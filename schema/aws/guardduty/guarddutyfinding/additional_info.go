package guarddutyfinding

type AdditionalInfo struct {
    Anomalies Anomalies `json:"anomalies,omitempty"`
    NewPolicy NewPolicy `json:"newPolicy,omitempty"`
    OldPolicy OldPolicy `json:"oldPolicy,omitempty"`
    ProfiledBehavior ProfiledBehavior `json:"profiledBehavior,omitempty"`
    UnusualBehavior UnusualBehavior `json:"unusualBehavior,omitempty"`
    UserAgent UserAgent `json:"userAgent,omitempty"`
    AdditionalScannedPorts []interface{} `json:"additionalScannedPorts,omitempty"`
    ApiCalls []AdditionalInfoItem `json:"apiCalls,omitempty"`
    Domain string `json:"domain,omitempty"`
    InBytes string `json:"inBytes,omitempty"`
    LocalPort string `json:"localPort,omitempty"`
    OutBytes string `json:"outBytes,omitempty"`
    PortsScannedSample []float64 `json:"portsScannedSample,omitempty"`
    RecentCredentials []AdditionalInfoItem_1 `json:"recentCredentials,omitempty"`
    Sample bool `json:"sample"`
    ScannedPort float64 `json:"scannedPort,omitempty"`
    ThreatListName string `json:"threatListName,omitempty"`
    ThreatName string `json:"threatName,omitempty"`
    AdditionalInfoType string `json:"type"`
    Unusual OneOfAdditionalInfoUnusual `json:"unusual,omitempty"`
    UnusualProtocol string `json:"unusualProtocol,omitempty"`
    Value string `json:"value"`
}

func (a *AdditionalInfo) SetAnomalies(anomalies Anomalies) {
    a.Anomalies = anomalies
}

func (a *AdditionalInfo) SetNewPolicy(newPolicy NewPolicy) {
    a.NewPolicy = newPolicy
}

func (a *AdditionalInfo) SetOldPolicy(oldPolicy OldPolicy) {
    a.OldPolicy = oldPolicy
}

func (a *AdditionalInfo) SetProfiledBehavior(profiledBehavior ProfiledBehavior) {
    a.ProfiledBehavior = profiledBehavior
}

func (a *AdditionalInfo) SetUnusualBehavior(unusualBehavior UnusualBehavior) {
    a.UnusualBehavior = unusualBehavior
}

func (a *AdditionalInfo) SetUserAgent(userAgent UserAgent) {
    a.UserAgent = userAgent
}

func (a *AdditionalInfo) SetAdditionalScannedPorts(additionalScannedPorts []interface{}) {
    a.AdditionalScannedPorts = additionalScannedPorts
}

func (a *AdditionalInfo) SetApiCalls(apiCalls []AdditionalInfoItem) {
    a.ApiCalls = apiCalls
}

func (a *AdditionalInfo) SetDomain(domain string) {
    a.Domain = domain
}

func (a *AdditionalInfo) SetInBytes(inBytes string) {
    a.InBytes = inBytes
}

func (a *AdditionalInfo) SetLocalPort(localPort string) {
    a.LocalPort = localPort
}

func (a *AdditionalInfo) SetOutBytes(outBytes string) {
    a.OutBytes = outBytes
}

func (a *AdditionalInfo) SetPortsScannedSample(portsScannedSample []float64) {
    a.PortsScannedSample = portsScannedSample
}

func (a *AdditionalInfo) SetRecentCredentials(recentCredentials []AdditionalInfoItem_1) {
    a.RecentCredentials = recentCredentials
}

func (a *AdditionalInfo) SetSample(sample bool) {
    a.Sample = sample
}

func (a *AdditionalInfo) SetScannedPort(scannedPort float64) {
    a.ScannedPort = scannedPort
}

func (a *AdditionalInfo) SetThreatListName(threatListName string) {
    a.ThreatListName = threatListName
}

func (a *AdditionalInfo) SetThreatName(threatName string) {
    a.ThreatName = threatName
}

func (a *AdditionalInfo) SetAdditionalInfoType(additionalInfoType string) {
    a.AdditionalInfoType = additionalInfoType
}

func (a *AdditionalInfo) SetUnusual(unusual OneOfAdditionalInfoUnusual) {
    a.Unusual = unusual
}

func (a *AdditionalInfo) SetUnusualProtocol(unusualProtocol string) {
    a.UnusualProtocol = unusualProtocol
}

func (a *AdditionalInfo) SetValue(value string) {
    a.Value = value
}
