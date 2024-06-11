package guarddutyfinding

type UnusualBehavior struct {
    IsUnusualUserIdentity string `json:"isUnusualUserIdentity,omitempty"`
    NumberOfPast24HoursAPIsBucketProfiling string `json:"numberOfPast24HoursAPIsBucketProfiling,omitempty"`
    NumberOfPast24HoursAPIsBucketUserIdentityProfiling string `json:"numberOfPast24HoursAPIsBucketUserIdentityProfiling,omitempty"`
    NumberOfPast24HoursAPIsUserIdentityProfiling string `json:"numberOfPast24HoursAPIsUserIdentityProfiling,omitempty"`
    UnusualAPIsAccountProfiling string `json:"unusualAPIsAccountProfiling,omitempty"`
    UnusualAPIsUserIdentityProfiling string `json:"unusualAPIsUserIdentityProfiling,omitempty"`
    UnusualASNsAccountProfiling string `json:"unusualASNsAccountProfiling,omitempty"`
    UnusualASNsBucketProfiling string `json:"unusualASNsBucketProfiling,omitempty"`
    UnusualASNsUserIdentityProfiling string `json:"unusualASNsUserIdentityProfiling,omitempty"`
    UnusualBucketsAccountProfiling string `json:"unusualBucketsAccountProfiling,omitempty"`
    UnusualBucketsUserIdentityProfiling string `json:"unusualBucketsUserIdentityProfiling,omitempty"`
    UnusualUserAgentsAccountProfiling string `json:"unusualUserAgentsAccountProfiling,omitempty"`
    UnusualUserAgentsUserIdentityProfiling string `json:"unusualUserAgentsUserIdentityProfiling,omitempty"`
    UnusualUserNamesAccountProfiling string `json:"unusualUserNamesAccountProfiling,omitempty"`
    UnusualUserNamesBucketProfiling string `json:"unusualUserNamesBucketProfiling,omitempty"`
    UnusualUserTypesAccountProfiling string `json:"unusualUserTypesAccountProfiling,omitempty"`
}

func (u *UnusualBehavior) SetIsUnusualUserIdentity(isUnusualUserIdentity string) {
    u.IsUnusualUserIdentity = isUnusualUserIdentity
}

func (u *UnusualBehavior) SetNumberOfPast24HoursAPIsBucketProfiling(numberOfPast24HoursAPIsBucketProfiling string) {
    u.NumberOfPast24HoursAPIsBucketProfiling = numberOfPast24HoursAPIsBucketProfiling
}

func (u *UnusualBehavior) SetNumberOfPast24HoursAPIsBucketUserIdentityProfiling(numberOfPast24HoursAPIsBucketUserIdentityProfiling string) {
    u.NumberOfPast24HoursAPIsBucketUserIdentityProfiling = numberOfPast24HoursAPIsBucketUserIdentityProfiling
}

func (u *UnusualBehavior) SetNumberOfPast24HoursAPIsUserIdentityProfiling(numberOfPast24HoursAPIsUserIdentityProfiling string) {
    u.NumberOfPast24HoursAPIsUserIdentityProfiling = numberOfPast24HoursAPIsUserIdentityProfiling
}

func (u *UnusualBehavior) SetUnusualAPIsAccountProfiling(unusualAPIsAccountProfiling string) {
    u.UnusualAPIsAccountProfiling = unusualAPIsAccountProfiling
}

func (u *UnusualBehavior) SetUnusualAPIsUserIdentityProfiling(unusualAPIsUserIdentityProfiling string) {
    u.UnusualAPIsUserIdentityProfiling = unusualAPIsUserIdentityProfiling
}

func (u *UnusualBehavior) SetUnusualASNsAccountProfiling(unusualASNsAccountProfiling string) {
    u.UnusualASNsAccountProfiling = unusualASNsAccountProfiling
}

func (u *UnusualBehavior) SetUnusualASNsBucketProfiling(unusualASNsBucketProfiling string) {
    u.UnusualASNsBucketProfiling = unusualASNsBucketProfiling
}

func (u *UnusualBehavior) SetUnusualASNsUserIdentityProfiling(unusualASNsUserIdentityProfiling string) {
    u.UnusualASNsUserIdentityProfiling = unusualASNsUserIdentityProfiling
}

func (u *UnusualBehavior) SetUnusualBucketsAccountProfiling(unusualBucketsAccountProfiling string) {
    u.UnusualBucketsAccountProfiling = unusualBucketsAccountProfiling
}

func (u *UnusualBehavior) SetUnusualBucketsUserIdentityProfiling(unusualBucketsUserIdentityProfiling string) {
    u.UnusualBucketsUserIdentityProfiling = unusualBucketsUserIdentityProfiling
}

func (u *UnusualBehavior) SetUnusualUserAgentsAccountProfiling(unusualUserAgentsAccountProfiling string) {
    u.UnusualUserAgentsAccountProfiling = unusualUserAgentsAccountProfiling
}

func (u *UnusualBehavior) SetUnusualUserAgentsUserIdentityProfiling(unusualUserAgentsUserIdentityProfiling string) {
    u.UnusualUserAgentsUserIdentityProfiling = unusualUserAgentsUserIdentityProfiling
}

func (u *UnusualBehavior) SetUnusualUserNamesAccountProfiling(unusualUserNamesAccountProfiling string) {
    u.UnusualUserNamesAccountProfiling = unusualUserNamesAccountProfiling
}

func (u *UnusualBehavior) SetUnusualUserNamesBucketProfiling(unusualUserNamesBucketProfiling string) {
    u.UnusualUserNamesBucketProfiling = unusualUserNamesBucketProfiling
}

func (u *UnusualBehavior) SetUnusualUserTypesAccountProfiling(unusualUserTypesAccountProfiling string) {
    u.UnusualUserTypesAccountProfiling = unusualUserTypesAccountProfiling
}
