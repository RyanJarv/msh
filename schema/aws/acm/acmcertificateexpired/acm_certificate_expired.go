package acmcertificateexpired

import (
    "time"
)


type ACMCertificateExpired struct {
    CertificateType string `json:"CertificateType"`
    CommonName string `json:"CommonName"`
    DomainValidationMethod string `json:"DomainValidationMethod,omitempty"`
    CertificateCreatedDate time.Time `json:"CertificateCreatedDate"`
    CertificateExpirationDate time.Time `json:"CertificateExpirationDate"`
    InUse bool `json:"InUse"`
    Exported bool `json:"Exported"`
}

func (a *ACMCertificateExpired) SetCertificateType(certificateType string) {
    a.CertificateType = certificateType
}

func (a *ACMCertificateExpired) SetCommonName(commonName string) {
    a.CommonName = commonName
}

func (a *ACMCertificateExpired) SetDomainValidationMethod(domainValidationMethod string) {
    a.DomainValidationMethod = domainValidationMethod
}

func (a *ACMCertificateExpired) SetCertificateCreatedDate(certificateCreatedDate time.Time) {
    a.CertificateCreatedDate = certificateCreatedDate
}

func (a *ACMCertificateExpired) SetCertificateExpirationDate(certificateExpirationDate time.Time) {
    a.CertificateExpirationDate = certificateExpirationDate
}

func (a *ACMCertificateExpired) SetInUse(inUse bool) {
    a.InUse = inUse
}

func (a *ACMCertificateExpired) SetExported(exported bool) {
    a.Exported = exported
}
