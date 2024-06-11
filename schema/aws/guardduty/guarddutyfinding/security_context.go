package guarddutyfinding

type SecurityContext struct {
    Privileged bool `json:"privileged"`
}

func (s *SecurityContext) SetPrivileged(privileged bool) {
    s.Privileged = privileged
}
