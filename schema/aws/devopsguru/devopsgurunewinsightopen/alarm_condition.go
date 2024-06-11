package devopsgurunewinsightopen

type AlarmCondition struct {
    DetectionBand string `json:"detectionBand,omitempty"`
    ReferenceValue ReferenceValue `json:"referenceValue,omitempty"`
}

func (a *AlarmCondition) SetDetectionBand(detectionBand string) {
    a.DetectionBand = detectionBand
}

func (a *AlarmCondition) SetReferenceValue(referenceValue ReferenceValue) {
    a.ReferenceValue = referenceValue
}
