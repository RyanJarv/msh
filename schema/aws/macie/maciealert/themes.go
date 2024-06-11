package maciealert

type Themes struct {
    SecretMarkings float64 `json:"Secret Markings,omitempty"`
    ConfidentialMarkings float64 `json:"Confidential Markings,omitempty"`
    CorporateProposals float64 `json:"Corporate Proposals,omitempty"`
}

func (t *Themes) SetSecretMarkings(secretMarkings float64) {
    t.SecretMarkings = secretMarkings
}

func (t *Themes) SetConfidentialMarkings(confidentialMarkings float64) {
    t.ConfidentialMarkings = confidentialMarkings
}

func (t *Themes) SetCorporateProposals(corporateProposals float64) {
    t.CorporateProposals = corporateProposals
}
