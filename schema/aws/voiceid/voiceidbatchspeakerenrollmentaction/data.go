package voiceidbatchspeakerenrollmentaction

type Data struct {
    EnrollmentConfig EnrollmentConfig `json:"enrollmentConfig,omitempty"`
    InputDataConfig InputDataConfig `json:"inputDataConfig,omitempty"`
    OutputDataConfig OutputDataConfig `json:"outputDataConfig,omitempty"`
    DataAccessRoleArn string `json:"dataAccessRoleArn,omitempty"`
}

func (d *Data) SetEnrollmentConfig(enrollmentConfig EnrollmentConfig) {
    d.EnrollmentConfig = enrollmentConfig
}

func (d *Data) SetInputDataConfig(inputDataConfig InputDataConfig) {
    d.InputDataConfig = inputDataConfig
}

func (d *Data) SetOutputDataConfig(outputDataConfig OutputDataConfig) {
    d.OutputDataConfig = outputDataConfig
}

func (d *Data) SetDataAccessRoleArn(dataAccessRoleArn string) {
    d.DataAccessRoleArn = dataAccessRoleArn
}
