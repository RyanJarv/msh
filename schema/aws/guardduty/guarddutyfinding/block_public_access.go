package guarddutyfinding

type BlockPublicAccess struct {
    BlockPublicAcls bool `json:"blockPublicAcls,omitempty"`
    BlockPublicPolicy bool `json:"blockPublicPolicy,omitempty"`
    IgnorePublicAcls bool `json:"ignorePublicAcls,omitempty"`
    RestrictPublicBuckets bool `json:"restrictPublicBuckets,omitempty"`
}

func (b *BlockPublicAccess) SetBlockPublicAcls(blockPublicAcls bool) {
    b.BlockPublicAcls = blockPublicAcls
}

func (b *BlockPublicAccess) SetBlockPublicPolicy(blockPublicPolicy bool) {
    b.BlockPublicPolicy = blockPublicPolicy
}

func (b *BlockPublicAccess) SetIgnorePublicAcls(ignorePublicAcls bool) {
    b.IgnorePublicAcls = ignorePublicAcls
}

func (b *BlockPublicAccess) SetRestrictPublicBuckets(restrictPublicBuckets bool) {
    b.RestrictPublicBuckets = restrictPublicBuckets
}
