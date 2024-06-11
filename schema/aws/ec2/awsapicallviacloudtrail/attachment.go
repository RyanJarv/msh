package awsapicallviacloudtrail

type Attachment struct {
    AttachmentId string `json:"attachmentId"`
    DeleteOnTermination bool `json:"deleteOnTermination"`
    DeviceIndex float64 `json:"deviceIndex"`
    AttachTime int64 `json:"attachTime"`
    Status string `json:"status"`
}

func (a *Attachment) SetAttachmentId(attachmentId string) {
    a.AttachmentId = attachmentId
}

func (a *Attachment) SetDeleteOnTermination(deleteOnTermination bool) {
    a.DeleteOnTermination = deleteOnTermination
}

func (a *Attachment) SetDeviceIndex(deviceIndex float64) {
    a.DeviceIndex = deviceIndex
}

func (a *Attachment) SetAttachTime(attachTime int64) {
    a.AttachTime = attachTime
}

func (a *Attachment) SetStatus(status string) {
    a.Status = status
}
