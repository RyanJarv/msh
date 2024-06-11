package maciealert

type Events struct {
    GetBucketLocation GetBucketLocation `json:"GetBucketLocation,omitempty"`
    GetBucketAcl GetBucketLocation `json:"GetBucketAcl,omitempty"`
    ListBuckets GetBucketLocation `json:"ListBuckets,omitempty"`
    GetBucketPolicy GetBucketPolicy `json:"GetBucketPolicy,omitempty"`
    ListRoles GetBucketLocation `json:"ListRoles,omitempty"`
}

func (e *Events) SetGetBucketLocation(getBucketLocation GetBucketLocation) {
    e.GetBucketLocation = getBucketLocation
}

func (e *Events) SetGetBucketAcl(getBucketAcl GetBucketLocation) {
    e.GetBucketAcl = getBucketAcl
}

func (e *Events) SetListBuckets(listBuckets GetBucketLocation) {
    e.ListBuckets = listBuckets
}

func (e *Events) SetGetBucketPolicy(getBucketPolicy GetBucketPolicy) {
    e.GetBucketPolicy = getBucketPolicy
}

func (e *Events) SetListRoles(listRoles GetBucketLocation) {
    e.ListRoles = listRoles
}
