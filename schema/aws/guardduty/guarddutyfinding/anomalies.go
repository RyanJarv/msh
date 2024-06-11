package guarddutyfinding

type Anomalies struct {
    AnomalousAPIs string `json:"anomalousAPIs,omitempty"`
}

func (a *Anomalies) SetAnomalousAPIs(anomalousAPIs string) {
    a.AnomalousAPIs = anomalousAPIs
}
