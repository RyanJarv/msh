package awsapicallviacloudtrail

type GroupSet_1Item struct {
    GroupId string `json:"groupId"`
}

func (g *GroupSet_1Item) SetGroupId(groupId string) {
    g.GroupId = groupId
}
