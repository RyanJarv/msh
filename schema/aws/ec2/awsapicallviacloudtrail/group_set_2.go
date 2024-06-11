package awsapicallviacloudtrail

type GroupSet_2 struct {
    Items []GroupSet_2Item `json:"items,omitempty"`
}

func (g *GroupSet_2) SetItems(items []GroupSet_2Item) {
    g.Items = items
}
