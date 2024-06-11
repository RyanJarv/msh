package applicationcostprofileringestiondataaccessfailure

type ApplicationCostProfilerIngestionDataAccessFailure struct {
    Message string `json:"message"`
}

func (a *ApplicationCostProfilerIngestionDataAccessFailure) SetMessage(message string) {
    a.Message = message
}
