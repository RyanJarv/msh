package guarddutyfinding

type Organization_1 struct {
    Asn string `json:"asn,omitempty"`
    AsnOrg string `json:"asnOrg,omitempty"`
    Isp string `json:"isp,omitempty"`
    Org string `json:"org,omitempty"`
}

func (o *Organization_1) SetAsn(asn string) {
    o.Asn = asn
}

func (o *Organization_1) SetAsnOrg(asnOrg string) {
    o.AsnOrg = asnOrg
}

func (o *Organization_1) SetIsp(isp string) {
    o.Isp = isp
}

func (o *Organization_1) SetOrg(org string) {
    o.Org = org
}
