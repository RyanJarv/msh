package trustedadvisorcheckitemrefreshnotification

type TrustedAdvisorCheckItemRefreshNotification struct {
    CheckItemDetail Check_item_detail `json:"check-item-detail"`
    CheckName string `json:"check-name"`
    ResourceId string `json:"resource_id"`
    Status string `json:"status"`
    Uuid string `json:"uuid"`
}

func (t *TrustedAdvisorCheckItemRefreshNotification) SetCheckItemDetail(checkItemDetail Check_item_detail) {
    t.CheckItemDetail = checkItemDetail
}

func (t *TrustedAdvisorCheckItemRefreshNotification) SetCheckName(checkName string) {
    t.CheckName = checkName
}

func (t *TrustedAdvisorCheckItemRefreshNotification) SetResourceId(resourceId string) {
    t.ResourceId = resourceId
}

func (t *TrustedAdvisorCheckItemRefreshNotification) SetStatus(status string) {
    t.Status = status
}

func (t *TrustedAdvisorCheckItemRefreshNotification) SetUuid(uuid string) {
    t.Uuid = uuid
}
