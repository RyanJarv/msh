package voiceidfraudsteraction

type Data struct {
    RegistrationSource string `json:"registrationSource,omitempty"`
    RegistrationSourceId string `json:"registrationSourceId,omitempty"`
    RegistrationStatus string `json:"registrationStatus,omitempty"`
}

func (d *Data) SetRegistrationSource(registrationSource string) {
    d.RegistrationSource = registrationSource
}

func (d *Data) SetRegistrationSourceId(registrationSourceId string) {
    d.RegistrationSourceId = registrationSourceId
}

func (d *Data) SetRegistrationStatus(registrationStatus string) {
    d.RegistrationStatus = registrationStatus
}
