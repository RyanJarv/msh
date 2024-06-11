package guarddutyfinding

type PermissionConfiguration struct {
    AccountLevelPermissions AccountLevelPermissions `json:"accountLevelPermissions,omitempty"`
    BucketLevelPermissions BucketLevelPermissions `json:"bucketLevelPermissions,omitempty"`
}

func (p *PermissionConfiguration) SetAccountLevelPermissions(accountLevelPermissions AccountLevelPermissions) {
    p.AccountLevelPermissions = accountLevelPermissions
}

func (p *PermissionConfiguration) SetBucketLevelPermissions(bucketLevelPermissions BucketLevelPermissions) {
    p.BucketLevelPermissions = bucketLevelPermissions
}
