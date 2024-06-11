package awsapicallviacloudtrail

import (
    "time"
)


type Attributes struct {
    CreationDate time.Time `json:"creationDate,omitempty"`
    MfaAuthenticated string `json:"mfaAuthenticated,omitempty"`
}

func (a *Attributes) SetCreationDate(creationDate time.Time) {
    a.CreationDate = creationDate
}

func (a *Attributes) SetMfaAuthenticated(mfaAuthenticated string) {
    a.MfaAuthenticated = mfaAuthenticated
}
