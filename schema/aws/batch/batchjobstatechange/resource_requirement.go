package batchjobstatechange

type ResourceRequirement struct {
    ResourceRequirementType string `json:"type,omitempty"`
    Value string `json:"value,omitempty"`
}

func (r *ResourceRequirement) SetResourceRequirementType(resourceRequirementType string) {
    r.ResourceRequirementType = resourceRequirementType
}

func (r *ResourceRequirement) SetValue(value string) {
    r.Value = value
}
