package guarddutyfinding

type Organization struct {
    Asn string `json:"asn,omitempty"`
    AsnOrg string `json:"asnOrg,omitempty"`
    Isp string `json:"isp,omitempty"`
    Org string `json:"org,omitempty"`
}

func (o *Organization) SetAsn(asn string) {
    o.Asn = asn
}

func (o *Organization) SetAsnOrg(asnOrg string) {
    o.AsnOrg = asnOrg
}

func (o *Organization) SetIsp(isp string) {
    o.Isp = isp
}

func (o *Organization) SetOrg(org string) {
    o.Org = org
}
