package maciealert

type Owner_1 struct {
    DisplayName string `json:"DisplayName"`
    ID string `json:"ID"`
}

func (o *Owner_1) SetDisplayName(displayName string) {
    o.DisplayName = displayName
}

func (o *Owner_1) SetID(iD string) {
    o.ID = iD
}
