package maciealert

type GetBucketPolicy struct {
    ISP ISP `json:"ISP,omitempty"`
    ErrorCode Error_Code `json:"Error-Code,omitempty"`
    Count float64 `json:"count,omitempty"`
}

func (g *GetBucketPolicy) SetISP(iSP ISP) {
    g.ISP = iSP
}

func (g *GetBucketPolicy) SetErrorCode(errorCode Error_Code) {
    g.ErrorCode = errorCode
}

func (g *GetBucketPolicy) SetCount(count float64) {
    g.Count = count
}
