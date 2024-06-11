package awsapicallviacloudtrail

type GroupSet_1 struct {
    Items []GroupSet_1Item `json:"items,omitempty"`
}

func (g *GroupSet_1) SetItems(items []GroupSet_1Item) {
    g.Items = items
}
