package awsapicallviacloudtrail

type RequestParametersItem struct {
    ModelName string `json:"modelName"`
    InstanceType string `json:"instanceType"`
    InitialVariantWeight float64 `json:"initialVariantWeight"`
    VariantName string `json:"variantName"`
    InitialInstanceCount float64 `json:"initialInstanceCount"`
}

func (r *RequestParametersItem) SetModelName(modelName string) {
    r.ModelName = modelName
}

func (r *RequestParametersItem) SetInstanceType(instanceType string) {
    r.InstanceType = instanceType
}

func (r *RequestParametersItem) SetInitialVariantWeight(initialVariantWeight float64) {
    r.InitialVariantWeight = initialVariantWeight
}

func (r *RequestParametersItem) SetVariantName(variantName string) {
    r.VariantName = variantName
}

func (r *RequestParametersItem) SetInitialInstanceCount(initialInstanceCount float64) {
    r.InitialInstanceCount = initialInstanceCount
}
