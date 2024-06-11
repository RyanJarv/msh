package trustedadvisorcheckitemrefreshnotification

type Check_item_detail struct {
    EC2ConfigStatus string `json:"EC2Config Status"`
    InstanceID string `json:"Instance ID"`
    Region string `json:"Region"`
    Status string `json:"Status"`
    Timestamp string `json:"Timestamp"`
}

func (c *Check_item_detail) SetEC2ConfigStatus(eC2ConfigStatus string) {
    c.EC2ConfigStatus = eC2ConfigStatus
}

func (c *Check_item_detail) SetInstanceID(instanceID string) {
    c.InstanceID = instanceID
}

func (c *Check_item_detail) SetRegion(region string) {
    c.Region = region
}

func (c *Check_item_detail) SetStatus(status string) {
    c.Status = status
}

func (c *Check_item_detail) SetTimestamp(timestamp string) {
    c.Timestamp = timestamp
}
