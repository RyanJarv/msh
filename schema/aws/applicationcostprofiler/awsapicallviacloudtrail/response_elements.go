package awsapicallviacloudtrail

type ResponseElements struct {
    ImportId string `json:"importId,omitempty"`
    Message string `json:"message,omitempty"`
    ReportId string `json:"reportId,omitempty"`
}

func (r *ResponseElements) SetImportId(importId string) {
    r.ImportId = importId
}

func (r *ResponseElements) SetMessage(message string) {
    r.Message = message
}

func (r *ResponseElements) SetReportId(reportId string) {
    r.ReportId = reportId
}
