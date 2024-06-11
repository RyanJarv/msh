package guarddutyfinding

type AccessControlList struct {
    AllowsPublicReadAccess bool `json:"allowsPublicReadAccess,omitempty"`
    AllowsPublicWriteAccess bool `json:"allowsPublicWriteAccess,omitempty"`
}

func (a *AccessControlList) SetAllowsPublicReadAccess(allowsPublicReadAccess bool) {
    a.AllowsPublicReadAccess = allowsPublicReadAccess
}

func (a *AccessControlList) SetAllowsPublicWriteAccess(allowsPublicWriteAccess bool) {
    a.AllowsPublicWriteAccess = allowsPublicWriteAccess
}
