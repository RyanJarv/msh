package maciealert

type ACLItem struct {
    Owner Owner_1 `json:"Owner"`
    Grants []ACLItemItem `json:"Grants"`
}

func (a *ACLItem) SetOwner(owner Owner_1) {
    a.Owner = owner
}

func (a *ACLItem) SetGrants(grants []ACLItemItem) {
    a.Grants = grants
}
