package applicationcostprofileringestiondatavalidationfailure

type ApplicationCostProfilerIngestionDataValidationFailure struct {
    Message string `json:"message"`
}

func (a *ApplicationCostProfilerIngestionDataValidationFailure) SetMessage(message string) {
    a.Message = message
}
