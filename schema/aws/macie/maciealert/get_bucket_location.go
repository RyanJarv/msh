package maciealert

type GetBucketLocation struct {
    ISP ISP `json:"ISP,omitempty"`
    Count float64 `json:"count,omitempty"`
}

func (g *GetBucketLocation) SetISP(iSP ISP) {
    g.ISP = iSP
}

func (g *GetBucketLocation) SetCount(count float64) {
    g.Count = count
}
