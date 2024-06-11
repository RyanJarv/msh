package guarddutyfinding

type Organization_2 struct {
    Asn string `json:"asn,omitempty"`
    AsnOrg string `json:"asnOrg,omitempty"`
    Isp string `json:"isp,omitempty"`
    Org string `json:"org,omitempty"`
}

func (o *Organization_2) SetAsn(asn string) {
    o.Asn = asn
}

func (o *Organization_2) SetAsnOrg(asnOrg string) {
    o.AsnOrg = asnOrg
}

func (o *Organization_2) SetIsp(isp string) {
    o.Isp = isp
}

func (o *Organization_2) SetOrg(org string) {
    o.Org = org
}
