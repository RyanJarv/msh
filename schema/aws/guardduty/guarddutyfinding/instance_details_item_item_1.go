package guarddutyfinding

type InstanceDetailsItemItem_1 struct {
    GroupId string `json:"groupId"`
    GroupName string `json:"groupName"`
}

func (i *InstanceDetailsItemItem_1) SetGroupId(groupId string) {
    i.GroupId = groupId
}

func (i *InstanceDetailsItemItem_1) SetGroupName(groupName string) {
    i.GroupName = groupName
}
