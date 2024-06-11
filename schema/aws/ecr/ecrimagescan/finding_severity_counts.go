package ecrimagescan

type FindingSeverityCounts struct {
    CRITICAL float64 `json:"CRITICAL,omitempty"`
    HIGH float64 `json:"HIGH,omitempty"`
    INFORMATIONAL float64 `json:"INFORMATIONAL,omitempty"`
    LOW float64 `json:"LOW,omitempty"`
    MEDIUM float64 `json:"MEDIUM,omitempty"`
    UNDEFINED float64 `json:"UNDEFINED,omitempty"`
}

func (f *FindingSeverityCounts) SetCRITICAL(cRITICAL float64) {
    f.CRITICAL = cRITICAL
}

func (f *FindingSeverityCounts) SetHIGH(hIGH float64) {
    f.HIGH = hIGH
}

func (f *FindingSeverityCounts) SetINFORMATIONAL(iNFORMATIONAL float64) {
    f.INFORMATIONAL = iNFORMATIONAL
}

func (f *FindingSeverityCounts) SetLOW(lOW float64) {
    f.LOW = lOW
}

func (f *FindingSeverityCounts) SetMEDIUM(mEDIUM float64) {
    f.MEDIUM = mEDIUM
}

func (f *FindingSeverityCounts) SetUNDEFINED(uNDEFINED float64) {
    f.UNDEFINED = uNDEFINED
}
