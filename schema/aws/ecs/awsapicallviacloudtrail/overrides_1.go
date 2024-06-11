package awsapicallviacloudtrail

type Overrides_1 struct {
    ContainerOverrides []Overrides_1Item `json:"containerOverrides"`
    InferenceAcceleratorOverrides []interface{} `json:"inferenceAcceleratorOverrides"`
}

func (o *Overrides_1) SetContainerOverrides(containerOverrides []Overrides_1Item) {
    o.ContainerOverrides = containerOverrides
}

func (o *Overrides_1) SetInferenceAcceleratorOverrides(inferenceAcceleratorOverrides []interface{}) {
    o.InferenceAcceleratorOverrides = inferenceAcceleratorOverrides
}
