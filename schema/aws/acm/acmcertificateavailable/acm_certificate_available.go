package acmcertificateavailable

import (
    "time"
)


type ACMCertificateAvailable struct {
    CertificateType string `json:"CertificateType"`
    CommonName string `json:"CommonName"`
    Action string `json:"Action"`
    DomainValidationMethod string `json:"DomainValidationMethod,omitempty"`
    DaysToExpiry float64 `json:"DaysToExpiry"`
    CertificateCreatedDate time.Time `json:"CertificateCreatedDate"`
    CertificateExpirationDate time.Time `json:"CertificateExpirationDate"`
    InUse bool `json:"InUse"`
    Exported bool `json:"Exported"`
}

func (a *ACMCertificateAvailable) SetCertificateType(certificateType string) {
    a.CertificateType = certificateType
}

func (a *ACMCertificateAvailable) SetCommonName(commonName string) {
    a.CommonName = commonName
}

func (a *ACMCertificateAvailable) SetAction(action string) {
    a.Action = action
}

func (a *ACMCertificateAvailable) SetDomainValidationMethod(domainValidationMethod string) {
    a.DomainValidationMethod = domainValidationMethod
}

func (a *ACMCertificateAvailable) SetDaysToExpiry(daysToExpiry float64) {
    a.DaysToExpiry = daysToExpiry
}

func (a *ACMCertificateAvailable) SetCertificateCreatedDate(certificateCreatedDate time.Time) {
    a.CertificateCreatedDate = certificateCreatedDate
}

func (a *ACMCertificateAvailable) SetCertificateExpirationDate(certificateExpirationDate time.Time) {
    a.CertificateExpirationDate = certificateExpirationDate
}

func (a *ACMCertificateAvailable) SetInUse(inUse bool) {
    a.InUse = inUse
}

func (a *ACMCertificateAvailable) SetExported(exported bool) {
    a.Exported = exported
}
