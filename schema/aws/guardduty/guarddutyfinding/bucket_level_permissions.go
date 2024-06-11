package guarddutyfinding

type BucketLevelPermissions struct {
    AccessControlList AccessControlList `json:"accessControlList,omitempty"`
    BlockPublicAccess BlockPublicAccess `json:"blockPublicAccess,omitempty"`
    BucketPolicy AccessControlList `json:"bucketPolicy,omitempty"`
}

func (b *BucketLevelPermissions) SetAccessControlList(accessControlList AccessControlList) {
    b.AccessControlList = accessControlList
}

func (b *BucketLevelPermissions) SetBlockPublicAccess(blockPublicAccess BlockPublicAccess) {
    b.BlockPublicAccess = blockPublicAccess
}

func (b *BucketLevelPermissions) SetBucketPolicy(bucketPolicy AccessControlList) {
    b.BucketPolicy = bucketPolicy
}
