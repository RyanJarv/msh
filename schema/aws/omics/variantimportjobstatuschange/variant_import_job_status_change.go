package variantimportjobstatuschange

type VariantImportJobStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    Arn string `json:"arn,omitempty"`
    StoreId string `json:"storeId,omitempty"`
    Status string `json:"status,omitempty"`
    StoreName string `json:"storeName,omitempty"`
    JobId string `json:"jobId,omitempty"`
}

func (v *VariantImportJobStatusChange) SetOmicsVersion(omicsVersion string) {
    v.OmicsVersion = omicsVersion
}

func (v *VariantImportJobStatusChange) SetArn(arn string) {
    v.Arn = arn
}

func (v *VariantImportJobStatusChange) SetStoreId(storeId string) {
    v.StoreId = storeId
}

func (v *VariantImportJobStatusChange) SetStatus(status string) {
    v.Status = status
}

func (v *VariantImportJobStatusChange) SetStoreName(storeName string) {
    v.StoreName = storeName
}

func (v *VariantImportJobStatusChange) SetJobId(jobId string) {
    v.JobId = jobId
}
