package awsapicallviacloudtrail

import (
    "time"
)


type Attributes struct {
    MfaAuthenticated string `json:"mfaAuthenticated,omitempty"`
    CreationDate time.Time `json:"creationDate,omitempty"`
}

func (a *Attributes) SetMfaAuthenticated(mfaAuthenticated string) {
    a.MfaAuthenticated = mfaAuthenticated
}

func (a *Attributes) SetCreationDate(creationDate time.Time) {
    a.CreationDate = creationDate
}
