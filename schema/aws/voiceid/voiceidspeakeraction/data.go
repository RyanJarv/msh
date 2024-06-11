package voiceidspeakeraction

type Data struct {
    EnrollmentSource string `json:"enrollmentSource"`
    EnrollmentSourceId string `json:"enrollmentSourceId"`
    EnrollmentStatus string `json:"enrollmentStatus"`
}

func (d *Data) SetEnrollmentSource(enrollmentSource string) {
    d.EnrollmentSource = enrollmentSource
}

func (d *Data) SetEnrollmentSourceId(enrollmentSourceId string) {
    d.EnrollmentSourceId = enrollmentSourceId
}

func (d *Data) SetEnrollmentStatus(enrollmentStatus string) {
    d.EnrollmentStatus = enrollmentStatus
}
