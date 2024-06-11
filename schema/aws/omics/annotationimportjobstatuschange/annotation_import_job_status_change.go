package annotationimportjobstatuschange

type AnnotationImportJobStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    Arn string `json:"arn,omitempty"`
    StoreId string `json:"storeId,omitempty"`
    Status string `json:"status,omitempty"`
    StoreName string `json:"storeName,omitempty"`
    JobId string `json:"jobId,omitempty"`
}

func (a *AnnotationImportJobStatusChange) SetOmicsVersion(omicsVersion string) {
    a.OmicsVersion = omicsVersion
}

func (a *AnnotationImportJobStatusChange) SetArn(arn string) {
    a.Arn = arn
}

func (a *AnnotationImportJobStatusChange) SetStoreId(storeId string) {
    a.StoreId = storeId
}

func (a *AnnotationImportJobStatusChange) SetStatus(status string) {
    a.Status = status
}

func (a *AnnotationImportJobStatusChange) SetStoreName(storeName string) {
    a.StoreName = storeName
}

func (a *AnnotationImportJobStatusChange) SetJobId(jobId string) {
    a.JobId = jobId
}
