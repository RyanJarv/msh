package applicationcostprofilerreportdeliveryfailure

type ApplicationCostProfilerReportDeliveryFailure struct {
    Message string `json:"message"`
}

func (a *ApplicationCostProfilerReportDeliveryFailure) SetMessage(message string) {
    a.Message = message
}
