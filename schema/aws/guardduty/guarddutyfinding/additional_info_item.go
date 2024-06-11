package guarddutyfinding

type AdditionalInfoItem struct {
    Count float64 `json:"count"`
    FirstSeen float64 `json:"firstSeen"`
    LastSeen float64 `json:"lastSeen"`
    Name string `json:"name"`
}

func (a *AdditionalInfoItem) SetCount(count float64) {
    a.Count = count
}

func (a *AdditionalInfoItem) SetFirstSeen(firstSeen float64) {
    a.FirstSeen = firstSeen
}

func (a *AdditionalInfoItem) SetLastSeen(lastSeen float64) {
    a.LastSeen = lastSeen
}

func (a *AdditionalInfoItem) SetName(name string) {
    a.Name = name
}
