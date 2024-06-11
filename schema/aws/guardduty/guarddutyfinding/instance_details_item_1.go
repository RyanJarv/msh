package guarddutyfinding

type InstanceDetailsItem_1 struct {
    ProductCodeId string `json:"productCodeId"`
    ProductCodeType string `json:"productCodeType"`
}

func (i *InstanceDetailsItem_1) SetProductCodeId(productCodeId string) {
    i.ProductCodeId = productCodeId
}

func (i *InstanceDetailsItem_1) SetProductCodeType(productCodeType string) {
    i.ProductCodeType = productCodeType
}
