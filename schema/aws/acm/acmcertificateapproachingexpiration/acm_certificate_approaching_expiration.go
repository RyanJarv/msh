package acmcertificateapproachingexpiration

type ACMCertificateApproachingExpiration struct {
    DaysToExpiry float64 `json:"DaysToExpiry"`
}

func (a *ACMCertificateApproachingExpiration) SetDaysToExpiry(daysToExpiry float64) {
    a.DaysToExpiry = daysToExpiry
}
