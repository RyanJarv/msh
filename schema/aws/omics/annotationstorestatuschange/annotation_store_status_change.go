package annotationstorestatuschange

type AnnotationStoreStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    Arn string `json:"arn,omitempty"`
    StoreId string `json:"storeId,omitempty"`
    Status string `json:"status,omitempty"`
    StoreName string `json:"storeName,omitempty"`
}

func (a *AnnotationStoreStatusChange) SetOmicsVersion(omicsVersion string) {
    a.OmicsVersion = omicsVersion
}

func (a *AnnotationStoreStatusChange) SetArn(arn string) {
    a.Arn = arn
}

func (a *AnnotationStoreStatusChange) SetStoreId(storeId string) {
    a.StoreId = storeId
}

func (a *AnnotationStoreStatusChange) SetStatus(status string) {
    a.Status = status
}

func (a *AnnotationStoreStatusChange) SetStoreName(storeName string) {
    a.StoreName = storeName
}
