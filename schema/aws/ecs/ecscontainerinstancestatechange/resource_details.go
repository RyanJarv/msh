package ecscontainerinstancestatechange

type ResourceDetails struct {
    IntegerValue float64 `json:"integerValue,omitempty"`
    LongValue float64 `json:"longValue,omitempty"`
    DoubleValue float64 `json:"doubleValue,omitempty"`
    Name string `json:"name,omitempty"`
    StringSetValue []string `json:"stringSetValue,omitempty"`
    ResourceDetailsType string `json:"type,omitempty"`
}

func (r *ResourceDetails) SetIntegerValue(integerValue float64) {
    r.IntegerValue = integerValue
}

func (r *ResourceDetails) SetLongValue(longValue float64) {
    r.LongValue = longValue
}

func (r *ResourceDetails) SetDoubleValue(doubleValue float64) {
    r.DoubleValue = doubleValue
}

func (r *ResourceDetails) SetName(name string) {
    r.Name = name
}

func (r *ResourceDetails) SetStringSetValue(stringSetValue []string) {
    r.StringSetValue = stringSetValue
}

func (r *ResourceDetails) SetResourceDetailsType(resourceDetailsType string) {
    r.ResourceDetailsType = resourceDetailsType
}
