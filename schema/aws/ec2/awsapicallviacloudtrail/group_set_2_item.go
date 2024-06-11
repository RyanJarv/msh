package awsapicallviacloudtrail

type GroupSet_2Item struct {
    GroupName string `json:"groupName"`
    GroupId string `json:"groupId"`
}

func (g *GroupSet_2Item) SetGroupName(groupName string) {
    g.GroupName = groupName
}

func (g *GroupSet_2Item) SetGroupId(groupId string) {
    g.GroupId = groupId
}
