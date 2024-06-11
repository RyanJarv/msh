package variantstorestatuschange

type VariantStoreStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    Arn string `json:"arn,omitempty"`
    StoreId string `json:"storeId,omitempty"`
    Status string `json:"status,omitempty"`
    StoreName string `json:"storeName,omitempty"`
}

func (v *VariantStoreStatusChange) SetOmicsVersion(omicsVersion string) {
    v.OmicsVersion = omicsVersion
}

func (v *VariantStoreStatusChange) SetArn(arn string) {
    v.Arn = arn
}

func (v *VariantStoreStatusChange) SetStoreId(storeId string) {
    v.StoreId = storeId
}

func (v *VariantStoreStatusChange) SetStatus(status string) {
    v.Status = status
}

func (v *VariantStoreStatusChange) SetStoreName(storeName string) {
    v.StoreName = storeName
}
