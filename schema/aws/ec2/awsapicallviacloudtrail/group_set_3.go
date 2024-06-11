package awsapicallviacloudtrail

type GroupSet_3 struct {
    Items []GroupSet_2Item `json:"items"`
}

func (g *GroupSet_3) SetItems(items []GroupSet_2Item) {
    g.Items = items
}
