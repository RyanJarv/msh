package guarddutyfinding

type Organization_4 struct {
    Asn string `json:"asn"`
    AsnOrg string `json:"asnOrg"`
    Isp string `json:"isp"`
    Org string `json:"org"`
}

func (o *Organization_4) SetAsn(asn string) {
    o.Asn = asn
}

func (o *Organization_4) SetAsnOrg(asnOrg string) {
    o.AsnOrg = asnOrg
}

func (o *Organization_4) SetIsp(isp string) {
    o.Isp = isp
}

func (o *Organization_4) SetOrg(org string) {
    o.Org = org
}
