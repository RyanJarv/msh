package guarddutyfinding

type ProfiledBehavior struct {
    FrequentProfiledAPIsAccountProfiling string `json:"frequentProfiledAPIsAccountProfiling,omitempty"`
    FrequentProfiledAPIsUserIdentityProfiling string `json:"frequentProfiledAPIsUserIdentityProfiling,omitempty"`
    FrequentProfiledASNsAccountProfiling string `json:"frequentProfiledASNsAccountProfiling,omitempty"`
    FrequentProfiledASNsBucketProfiling string `json:"frequentProfiledASNsBucketProfiling,omitempty"`
    FrequentProfiledASNsUserIdentityProfiling string `json:"frequentProfiledASNsUserIdentityProfiling,omitempty"`
    FrequentProfiledBucketsAccountProfiling string `json:"frequentProfiledBucketsAccountProfiling,omitempty"`
    FrequentProfiledBucketsUserIdentityProfiling string `json:"frequentProfiledBucketsUserIdentityProfiling,omitempty"`
    FrequentProfiledUserAgentsAccountProfiling string `json:"frequentProfiledUserAgentsAccountProfiling,omitempty"`
    FrequentProfiledUserAgentsUserIdentityProfiling string `json:"frequentProfiledUserAgentsUserIdentityProfiling,omitempty"`
    FrequentProfiledUserNamesAccountProfiling string `json:"frequentProfiledUserNamesAccountProfiling,omitempty"`
    FrequentProfiledUserNamesBucketProfiling string `json:"frequentProfiledUserNamesBucketProfiling,omitempty"`
    FrequentProfiledUserTypesAccountProfiling string `json:"frequentProfiledUserTypesAccountProfiling,omitempty"`
    InfrequentProfiledAPIsAccountProfiling string `json:"infrequentProfiledAPIsAccountProfiling,omitempty"`
    InfrequentProfiledAPIsUserIdentityProfiling string `json:"infrequentProfiledAPIsUserIdentityProfiling,omitempty"`
    InfrequentProfiledASNsAccountProfiling string `json:"infrequentProfiledASNsAccountProfiling,omitempty"`
    InfrequentProfiledASNsBucketProfiling string `json:"infrequentProfiledASNsBucketProfiling,omitempty"`
    InfrequentProfiledASNsUserIdentityProfiling string `json:"infrequentProfiledASNsUserIdentityProfiling,omitempty"`
    InfrequentProfiledBucketsAccountProfiling string `json:"infrequentProfiledBucketsAccountProfiling,omitempty"`
    InfrequentProfiledBucketsUserIdentityProfiling string `json:"infrequentProfiledBucketsUserIdentityProfiling,omitempty"`
    InfrequentProfiledUserAgentsAccountProfiling string `json:"infrequentProfiledUserAgentsAccountProfiling,omitempty"`
    InfrequentProfiledUserAgentsUserIdentityProfiling string `json:"infrequentProfiledUserAgentsUserIdentityProfiling,omitempty"`
    InfrequentProfiledUserNamesAccountProfiling string `json:"infrequentProfiledUserNamesAccountProfiling,omitempty"`
    InfrequentProfiledUserNamesBucketProfiling string `json:"infrequentProfiledUserNamesBucketProfiling,omitempty"`
    InfrequentProfiledUserTypesAccountProfiling string `json:"infrequentProfiledUserTypesAccountProfiling,omitempty"`
    NumberOfHistoricalDailyAvgAPIsBucketProfiling string `json:"numberOfHistoricalDailyAvgAPIsBucketProfiling,omitempty"`
    NumberOfHistoricalDailyAvgAPIsBucketUserIdentityProfiling string `json:"numberOfHistoricalDailyAvgAPIsBucketUserIdentityProfiling,omitempty"`
    NumberOfHistoricalDailyAvgAPIsUserIdentityProfiling string `json:"numberOfHistoricalDailyAvgAPIsUserIdentityProfiling,omitempty"`
    NumberOfHistoricalDailyMaxAPIsBucketProfiling string `json:"numberOfHistoricalDailyMaxAPIsBucketProfiling,omitempty"`
    NumberOfHistoricalDailyMaxAPIsBucketUserIdentityProfiling string `json:"numberOfHistoricalDailyMaxAPIsBucketUserIdentityProfiling,omitempty"`
    NumberOfHistoricalDailyMaxAPIsUserIdentityProfiling string `json:"numberOfHistoricalDailyMaxAPIsUserIdentityProfiling,omitempty"`
    RareProfiledAPIsAccountProfiling string `json:"rareProfiledAPIsAccountProfiling,omitempty"`
    RareProfiledAPIsUserIdentityProfiling string `json:"rareProfiledAPIsUserIdentityProfiling,omitempty"`
    RareProfiledASNsAccountProfiling string `json:"rareProfiledASNsAccountProfiling,omitempty"`
    RareProfiledASNsBucketProfiling string `json:"rareProfiledASNsBucketProfiling,omitempty"`
    RareProfiledASNsUserIdentityProfiling string `json:"rareProfiledASNsUserIdentityProfiling,omitempty"`
    RareProfiledBucketsAccountProfiling string `json:"rareProfiledBucketsAccountProfiling,omitempty"`
    RareProfiledBucketsUserIdentityProfiling string `json:"rareProfiledBucketsUserIdentityProfiling,omitempty"`
    RareProfiledUserAgentsAccountProfiling string `json:"rareProfiledUserAgentsAccountProfiling,omitempty"`
    RareProfiledUserAgentsUserIdentityProfiling string `json:"rareProfiledUserAgentsUserIdentityProfiling,omitempty"`
    RareProfiledUserNamesAccountProfiling string `json:"rareProfiledUserNamesAccountProfiling,omitempty"`
    RareProfiledUserNamesBucketProfiling string `json:"rareProfiledUserNamesBucketProfiling,omitempty"`
    RareProfiledUserTypesAccountProfiling string `json:"rareProfiledUserTypesAccountProfiling,omitempty"`
}

func (p *ProfiledBehavior) SetFrequentProfiledAPIsAccountProfiling(frequentProfiledAPIsAccountProfiling string) {
    p.FrequentProfiledAPIsAccountProfiling = frequentProfiledAPIsAccountProfiling
}

func (p *ProfiledBehavior) SetFrequentProfiledAPIsUserIdentityProfiling(frequentProfiledAPIsUserIdentityProfiling string) {
    p.FrequentProfiledAPIsUserIdentityProfiling = frequentProfiledAPIsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetFrequentProfiledASNsAccountProfiling(frequentProfiledASNsAccountProfiling string) {
    p.FrequentProfiledASNsAccountProfiling = frequentProfiledASNsAccountProfiling
}

func (p *ProfiledBehavior) SetFrequentProfiledASNsBucketProfiling(frequentProfiledASNsBucketProfiling string) {
    p.FrequentProfiledASNsBucketProfiling = frequentProfiledASNsBucketProfiling
}

func (p *ProfiledBehavior) SetFrequentProfiledASNsUserIdentityProfiling(frequentProfiledASNsUserIdentityProfiling string) {
    p.FrequentProfiledASNsUserIdentityProfiling = frequentProfiledASNsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetFrequentProfiledBucketsAccountProfiling(frequentProfiledBucketsAccountProfiling string) {
    p.FrequentProfiledBucketsAccountProfiling = frequentProfiledBucketsAccountProfiling
}

func (p *ProfiledBehavior) SetFrequentProfiledBucketsUserIdentityProfiling(frequentProfiledBucketsUserIdentityProfiling string) {
    p.FrequentProfiledBucketsUserIdentityProfiling = frequentProfiledBucketsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetFrequentProfiledUserAgentsAccountProfiling(frequentProfiledUserAgentsAccountProfiling string) {
    p.FrequentProfiledUserAgentsAccountProfiling = frequentProfiledUserAgentsAccountProfiling
}

func (p *ProfiledBehavior) SetFrequentProfiledUserAgentsUserIdentityProfiling(frequentProfiledUserAgentsUserIdentityProfiling string) {
    p.FrequentProfiledUserAgentsUserIdentityProfiling = frequentProfiledUserAgentsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetFrequentProfiledUserNamesAccountProfiling(frequentProfiledUserNamesAccountProfiling string) {
    p.FrequentProfiledUserNamesAccountProfiling = frequentProfiledUserNamesAccountProfiling
}

func (p *ProfiledBehavior) SetFrequentProfiledUserNamesBucketProfiling(frequentProfiledUserNamesBucketProfiling string) {
    p.FrequentProfiledUserNamesBucketProfiling = frequentProfiledUserNamesBucketProfiling
}

func (p *ProfiledBehavior) SetFrequentProfiledUserTypesAccountProfiling(frequentProfiledUserTypesAccountProfiling string) {
    p.FrequentProfiledUserTypesAccountProfiling = frequentProfiledUserTypesAccountProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledAPIsAccountProfiling(infrequentProfiledAPIsAccountProfiling string) {
    p.InfrequentProfiledAPIsAccountProfiling = infrequentProfiledAPIsAccountProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledAPIsUserIdentityProfiling(infrequentProfiledAPIsUserIdentityProfiling string) {
    p.InfrequentProfiledAPIsUserIdentityProfiling = infrequentProfiledAPIsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledASNsAccountProfiling(infrequentProfiledASNsAccountProfiling string) {
    p.InfrequentProfiledASNsAccountProfiling = infrequentProfiledASNsAccountProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledASNsBucketProfiling(infrequentProfiledASNsBucketProfiling string) {
    p.InfrequentProfiledASNsBucketProfiling = infrequentProfiledASNsBucketProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledASNsUserIdentityProfiling(infrequentProfiledASNsUserIdentityProfiling string) {
    p.InfrequentProfiledASNsUserIdentityProfiling = infrequentProfiledASNsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledBucketsAccountProfiling(infrequentProfiledBucketsAccountProfiling string) {
    p.InfrequentProfiledBucketsAccountProfiling = infrequentProfiledBucketsAccountProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledBucketsUserIdentityProfiling(infrequentProfiledBucketsUserIdentityProfiling string) {
    p.InfrequentProfiledBucketsUserIdentityProfiling = infrequentProfiledBucketsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledUserAgentsAccountProfiling(infrequentProfiledUserAgentsAccountProfiling string) {
    p.InfrequentProfiledUserAgentsAccountProfiling = infrequentProfiledUserAgentsAccountProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledUserAgentsUserIdentityProfiling(infrequentProfiledUserAgentsUserIdentityProfiling string) {
    p.InfrequentProfiledUserAgentsUserIdentityProfiling = infrequentProfiledUserAgentsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledUserNamesAccountProfiling(infrequentProfiledUserNamesAccountProfiling string) {
    p.InfrequentProfiledUserNamesAccountProfiling = infrequentProfiledUserNamesAccountProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledUserNamesBucketProfiling(infrequentProfiledUserNamesBucketProfiling string) {
    p.InfrequentProfiledUserNamesBucketProfiling = infrequentProfiledUserNamesBucketProfiling
}

func (p *ProfiledBehavior) SetInfrequentProfiledUserTypesAccountProfiling(infrequentProfiledUserTypesAccountProfiling string) {
    p.InfrequentProfiledUserTypesAccountProfiling = infrequentProfiledUserTypesAccountProfiling
}

func (p *ProfiledBehavior) SetNumberOfHistoricalDailyAvgAPIsBucketProfiling(numberOfHistoricalDailyAvgAPIsBucketProfiling string) {
    p.NumberOfHistoricalDailyAvgAPIsBucketProfiling = numberOfHistoricalDailyAvgAPIsBucketProfiling
}

func (p *ProfiledBehavior) SetNumberOfHistoricalDailyAvgAPIsBucketUserIdentityProfiling(numberOfHistoricalDailyAvgAPIsBucketUserIdentityProfiling string) {
    p.NumberOfHistoricalDailyAvgAPIsBucketUserIdentityProfiling = numberOfHistoricalDailyAvgAPIsBucketUserIdentityProfiling
}

func (p *ProfiledBehavior) SetNumberOfHistoricalDailyAvgAPIsUserIdentityProfiling(numberOfHistoricalDailyAvgAPIsUserIdentityProfiling string) {
    p.NumberOfHistoricalDailyAvgAPIsUserIdentityProfiling = numberOfHistoricalDailyAvgAPIsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetNumberOfHistoricalDailyMaxAPIsBucketProfiling(numberOfHistoricalDailyMaxAPIsBucketProfiling string) {
    p.NumberOfHistoricalDailyMaxAPIsBucketProfiling = numberOfHistoricalDailyMaxAPIsBucketProfiling
}

func (p *ProfiledBehavior) SetNumberOfHistoricalDailyMaxAPIsBucketUserIdentityProfiling(numberOfHistoricalDailyMaxAPIsBucketUserIdentityProfiling string) {
    p.NumberOfHistoricalDailyMaxAPIsBucketUserIdentityProfiling = numberOfHistoricalDailyMaxAPIsBucketUserIdentityProfiling
}

func (p *ProfiledBehavior) SetNumberOfHistoricalDailyMaxAPIsUserIdentityProfiling(numberOfHistoricalDailyMaxAPIsUserIdentityProfiling string) {
    p.NumberOfHistoricalDailyMaxAPIsUserIdentityProfiling = numberOfHistoricalDailyMaxAPIsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetRareProfiledAPIsAccountProfiling(rareProfiledAPIsAccountProfiling string) {
    p.RareProfiledAPIsAccountProfiling = rareProfiledAPIsAccountProfiling
}

func (p *ProfiledBehavior) SetRareProfiledAPIsUserIdentityProfiling(rareProfiledAPIsUserIdentityProfiling string) {
    p.RareProfiledAPIsUserIdentityProfiling = rareProfiledAPIsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetRareProfiledASNsAccountProfiling(rareProfiledASNsAccountProfiling string) {
    p.RareProfiledASNsAccountProfiling = rareProfiledASNsAccountProfiling
}

func (p *ProfiledBehavior) SetRareProfiledASNsBucketProfiling(rareProfiledASNsBucketProfiling string) {
    p.RareProfiledASNsBucketProfiling = rareProfiledASNsBucketProfiling
}

func (p *ProfiledBehavior) SetRareProfiledASNsUserIdentityProfiling(rareProfiledASNsUserIdentityProfiling string) {
    p.RareProfiledASNsUserIdentityProfiling = rareProfiledASNsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetRareProfiledBucketsAccountProfiling(rareProfiledBucketsAccountProfiling string) {
    p.RareProfiledBucketsAccountProfiling = rareProfiledBucketsAccountProfiling
}

func (p *ProfiledBehavior) SetRareProfiledBucketsUserIdentityProfiling(rareProfiledBucketsUserIdentityProfiling string) {
    p.RareProfiledBucketsUserIdentityProfiling = rareProfiledBucketsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetRareProfiledUserAgentsAccountProfiling(rareProfiledUserAgentsAccountProfiling string) {
    p.RareProfiledUserAgentsAccountProfiling = rareProfiledUserAgentsAccountProfiling
}

func (p *ProfiledBehavior) SetRareProfiledUserAgentsUserIdentityProfiling(rareProfiledUserAgentsUserIdentityProfiling string) {
    p.RareProfiledUserAgentsUserIdentityProfiling = rareProfiledUserAgentsUserIdentityProfiling
}

func (p *ProfiledBehavior) SetRareProfiledUserNamesAccountProfiling(rareProfiledUserNamesAccountProfiling string) {
    p.RareProfiledUserNamesAccountProfiling = rareProfiledUserNamesAccountProfiling
}

func (p *ProfiledBehavior) SetRareProfiledUserNamesBucketProfiling(rareProfiledUserNamesBucketProfiling string) {
    p.RareProfiledUserNamesBucketProfiling = rareProfiledUserNamesBucketProfiling
}

func (p *ProfiledBehavior) SetRareProfiledUserTypesAccountProfiling(rareProfiledUserTypesAccountProfiling string) {
    p.RareProfiledUserTypesAccountProfiling = rareProfiledUserTypesAccountProfiling
}
