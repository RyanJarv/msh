package awsapicallviacloudtrail

type FastRestoreRule struct {
    AvailabilityZones []string `json:"AvailabilityZones,omitempty"`
    Count float64 `json:"Count,omitempty"`
}

func (f *FastRestoreRule) SetAvailabilityZones(availabilityZones []string) {
    f.AvailabilityZones = availabilityZones
}

func (f *FastRestoreRule) SetCount(count float64) {
    f.Count = count
}
