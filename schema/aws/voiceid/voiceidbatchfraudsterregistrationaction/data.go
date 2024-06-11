package voiceidbatchfraudsterregistrationaction

type Data struct {
    InputDataConfig InputDataConfig `json:"inputDataConfig,omitempty"`
    OutputDataConfig OutputDataConfig `json:"outputDataConfig,omitempty"`
    RegistrationConfig RegistrationConfig `json:"registrationConfig,omitempty"`
    DataAccessRoleArn string `json:"dataAccessRoleArn,omitempty"`
}

func (d *Data) SetInputDataConfig(inputDataConfig InputDataConfig) {
    d.InputDataConfig = inputDataConfig
}

func (d *Data) SetOutputDataConfig(outputDataConfig OutputDataConfig) {
    d.OutputDataConfig = outputDataConfig
}

func (d *Data) SetRegistrationConfig(registrationConfig RegistrationConfig) {
    d.RegistrationConfig = registrationConfig
}

func (d *Data) SetDataAccessRoleArn(dataAccessRoleArn string) {
    d.DataAccessRoleArn = dataAccessRoleArn
}
