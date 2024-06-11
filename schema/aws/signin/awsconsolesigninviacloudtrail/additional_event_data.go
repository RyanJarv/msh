package awsconsolesigninviacloudtrail

type AdditionalEventData struct {
    LoginTo string `json:"LoginTo"`
    MobileVersion string `json:"MobileVersion"`
    MFAUsed string `json:"MFAUsed"`
}

func (a *AdditionalEventData) SetLoginTo(loginTo string) {
    a.LoginTo = loginTo
}

func (a *AdditionalEventData) SetMobileVersion(mobileVersion string) {
    a.MobileVersion = mobileVersion
}

func (a *AdditionalEventData) SetMFAUsed(mFAUsed string) {
    a.MFAUsed = mFAUsed
}
