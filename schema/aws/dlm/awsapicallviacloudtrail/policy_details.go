package awsapicallviacloudtrail

type PolicyDetails struct {
    PolicyType string `json:"PolicyType,omitempty"`
    ResourceTypes []string `json:"ResourceTypes,omitempty"`
    Schedules []PolicyDetailsItem `json:"Schedules,omitempty"`
    TargetTags []PolicyDetailsItem_1 `json:"TargetTags,omitempty"`
}

func (p *PolicyDetails) SetPolicyType(policyType string) {
    p.PolicyType = policyType
}

func (p *PolicyDetails) SetResourceTypes(resourceTypes []string) {
    p.ResourceTypes = resourceTypes
}

func (p *PolicyDetails) SetSchedules(schedules []PolicyDetailsItem) {
    p.Schedules = schedules
}

func (p *PolicyDetails) SetTargetTags(targetTags []PolicyDetailsItem_1) {
    p.TargetTags = targetTags
}
