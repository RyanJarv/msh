package guarddutyfinding

type ResourceItem struct {
    DefaultServerSideEncryption DefaultServerSideEncryption `json:"defaultServerSideEncryption,omitempty"`
    Owner Owner `json:"owner,omitempty"`
    PublicAccess PublicAccess `json:"publicAccess,omitempty"`
    Arn string `json:"arn,omitempty"`
    CreatedAt float64 `json:"createdAt,omitempty"`
    Name string `json:"name"`
    Tags []EcsClusterDetailsItem `json:"tags,omitempty"`
    ResourceItemType string `json:"type"`
}

func (r *ResourceItem) SetDefaultServerSideEncryption(defaultServerSideEncryption DefaultServerSideEncryption) {
    r.DefaultServerSideEncryption = defaultServerSideEncryption
}

func (r *ResourceItem) SetOwner(owner Owner) {
    r.Owner = owner
}

func (r *ResourceItem) SetPublicAccess(publicAccess PublicAccess) {
    r.PublicAccess = publicAccess
}

func (r *ResourceItem) SetArn(arn string) {
    r.Arn = arn
}

func (r *ResourceItem) SetCreatedAt(createdAt float64) {
    r.CreatedAt = createdAt
}

func (r *ResourceItem) SetName(name string) {
    r.Name = name
}

func (r *ResourceItem) SetTags(tags []EcsClusterDetailsItem) {
    r.Tags = tags
}

func (r *ResourceItem) SetResourceItemType(resourceItemType string) {
    r.ResourceItemType = resourceItemType
}
