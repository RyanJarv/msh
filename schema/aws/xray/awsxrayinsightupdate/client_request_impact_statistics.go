package awsxrayinsightupdate

type ClientRequestImpactStatistics struct {
    FaultCount float64 `json:"FaultCount"`
    OkCount float64 `json:"OkCount"`
    TotalCount float64 `json:"TotalCount"`
}

func (c *ClientRequestImpactStatistics) SetFaultCount(faultCount float64) {
    c.FaultCount = faultCount
}

func (c *ClientRequestImpactStatistics) SetOkCount(okCount float64) {
    c.OkCount = okCount
}

func (c *ClientRequestImpactStatistics) SetTotalCount(totalCount float64) {
    c.TotalCount = totalCount
}
