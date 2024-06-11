package acmcertificaterenewalactionrequired

import (
    "time"
)


type ACMCertificateRenewalActionRequired struct {
    CertificateType string `json:"CertificateType"`
    CommonName string `json:"CommonName"`
    DomainValidationMethod string `json:"DomainValidationMethod,omitempty"`
    RenewalStatusReason string `json:"RenewalStatusReason"`
    DaysToExpiry float64 `json:"DaysToExpiry"`
    CertificateCreatedDate time.Time `json:"CertificateCreatedDate"`
    CertificateExpirationDate time.Time `json:"CertificateExpirationDate"`
    InUse bool `json:"InUse"`
    Exported bool `json:"Exported"`
}

func (a *ACMCertificateRenewalActionRequired) SetCertificateType(certificateType string) {
    a.CertificateType = certificateType
}

func (a *ACMCertificateRenewalActionRequired) SetCommonName(commonName string) {
    a.CommonName = commonName
}

func (a *ACMCertificateRenewalActionRequired) SetDomainValidationMethod(domainValidationMethod string) {
    a.DomainValidationMethod = domainValidationMethod
}

func (a *ACMCertificateRenewalActionRequired) SetRenewalStatusReason(renewalStatusReason string) {
    a.RenewalStatusReason = renewalStatusReason
}

func (a *ACMCertificateRenewalActionRequired) SetDaysToExpiry(daysToExpiry float64) {
    a.DaysToExpiry = daysToExpiry
}

func (a *ACMCertificateRenewalActionRequired) SetCertificateCreatedDate(certificateCreatedDate time.Time) {
    a.CertificateCreatedDate = certificateCreatedDate
}

func (a *ACMCertificateRenewalActionRequired) SetCertificateExpirationDate(certificateExpirationDate time.Time) {
    a.CertificateExpirationDate = certificateExpirationDate
}

func (a *ACMCertificateRenewalActionRequired) SetInUse(inUse bool) {
    a.InUse = inUse
}

func (a *ACMCertificateRenewalActionRequired) SetExported(exported bool) {
    a.Exported = exported
}
