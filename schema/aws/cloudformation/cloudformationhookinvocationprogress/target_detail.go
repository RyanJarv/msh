package cloudformationhookinvocationprogress

type Target_detail struct {
    TargetAction string `json:"target-action"`
    TargetId string `json:"target-id"`
    TargetInvocationPoint string `json:"target-invocation-point"`
    TargetName string `json:"target-name"`
    TargetType string `json:"target-type"`
}

func (t *Target_detail) SetTargetAction(targetAction string) {
    t.TargetAction = targetAction
}

func (t *Target_detail) SetTargetId(targetId string) {
    t.TargetId = targetId
}

func (t *Target_detail) SetTargetInvocationPoint(targetInvocationPoint string) {
    t.TargetInvocationPoint = targetInvocationPoint
}

func (t *Target_detail) SetTargetName(targetName string) {
    t.TargetName = targetName
}

func (t *Target_detail) SetTargetType(targetType string) {
    t.TargetType = targetType
}
