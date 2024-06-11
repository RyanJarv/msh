package maciealert

type Grantee struct {
    GranteeType string `json:"Type"`
    URI string `json:"URI"`
}

func (g *Grantee) SetGranteeType(granteeType string) {
    g.GranteeType = granteeType
}

func (g *Grantee) SetURI(uRI string) {
    g.URI = uRI
}
