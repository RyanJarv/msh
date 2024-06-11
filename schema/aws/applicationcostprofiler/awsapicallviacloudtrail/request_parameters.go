package awsapicallviacloudtrail

type RequestParameters struct {
    DestinationS3Location DestinationS3Location `json:"destinationS3Location,omitempty"`
    SourceS3Location SourceS3Location `json:"sourceS3Location,omitempty"`
    Format string `json:"format,omitempty"`
    ReportDescription string `json:"reportDescription,omitempty"`
    ReportFrequency string `json:"reportFrequency,omitempty"`
    ReportId string `json:"reportId,omitempty"`
}

func (r *RequestParameters) SetDestinationS3Location(destinationS3Location DestinationS3Location) {
    r.DestinationS3Location = destinationS3Location
}

func (r *RequestParameters) SetSourceS3Location(sourceS3Location SourceS3Location) {
    r.SourceS3Location = sourceS3Location
}

func (r *RequestParameters) SetFormat(format string) {
    r.Format = format
}

func (r *RequestParameters) SetReportDescription(reportDescription string) {
    r.ReportDescription = reportDescription
}

func (r *RequestParameters) SetReportFrequency(reportFrequency string) {
    r.ReportFrequency = reportFrequency
}

func (r *RequestParameters) SetReportId(reportId string) {
    r.ReportId = reportId
}
