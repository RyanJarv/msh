package maciealert

type ACLItemItem struct {
    Grantee Grantee `json:"Grantee"`
    Permission string `json:"Permission"`
}

func (a *ACLItemItem) SetGrantee(grantee Grantee) {
    a.Grantee = grantee
}

func (a *ACLItemItem) SetPermission(permission string) {
    a.Permission = permission
}
