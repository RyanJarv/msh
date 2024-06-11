package awsapicallviacloudtrail

type RequestParametersItem struct {
    Address float64 `json:"address"`
    BigEndian bool `json:"bigEndian"`
    Capacity float64 `json:"capacity"`
    Hb []float64 `json:"hb"`
    IsReadOnly bool `json:"isReadOnly"`
    Limit float64 `json:"limit"`
    Mark float64 `json:"mark"`
    NativeByteOrder bool `json:"nativeByteOrder"`
    Offset float64 `json:"offset"`
    Position float64 `json:"position"`
}

func (r *RequestParametersItem) SetAddress(address float64) {
    r.Address = address
}

func (r *RequestParametersItem) SetBigEndian(bigEndian bool) {
    r.BigEndian = bigEndian
}

func (r *RequestParametersItem) SetCapacity(capacity float64) {
    r.Capacity = capacity
}

func (r *RequestParametersItem) SetHb(hb []float64) {
    r.Hb = hb
}

func (r *RequestParametersItem) SetIsReadOnly(isReadOnly bool) {
    r.IsReadOnly = isReadOnly
}

func (r *RequestParametersItem) SetLimit(limit float64) {
    r.Limit = limit
}

func (r *RequestParametersItem) SetMark(mark float64) {
    r.Mark = mark
}

func (r *RequestParametersItem) SetNativeByteOrder(nativeByteOrder bool) {
    r.NativeByteOrder = nativeByteOrder
}

func (r *RequestParametersItem) SetOffset(offset float64) {
    r.Offset = offset
}

func (r *RequestParametersItem) SetPosition(position float64) {
    r.Position = position
}
