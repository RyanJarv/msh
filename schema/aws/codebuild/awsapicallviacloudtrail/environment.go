package awsapicallviacloudtrail

type Environment struct {
    Certificate string `json:"certificate,omitempty"`
    ComputeType string `json:"computeType,omitempty"`
    EnvironmentVariables []EnvironmentItem `json:"environmentVariables,omitempty"`
    Image string `json:"image,omitempty"`
    ImagePullCredentialsType string `json:"imagePullCredentialsType,omitempty"`
    PrivilegedMode bool `json:"privilegedMode,omitempty"`
    EnvironmentType string `json:"type,omitempty"`
}

func (e *Environment) SetCertificate(certificate string) {
    e.Certificate = certificate
}

func (e *Environment) SetComputeType(computeType string) {
    e.ComputeType = computeType
}

func (e *Environment) SetEnvironmentVariables(environmentVariables []EnvironmentItem) {
    e.EnvironmentVariables = environmentVariables
}

func (e *Environment) SetImage(image string) {
    e.Image = image
}

func (e *Environment) SetImagePullCredentialsType(imagePullCredentialsType string) {
    e.ImagePullCredentialsType = imagePullCredentialsType
}

func (e *Environment) SetPrivilegedMode(privilegedMode bool) {
    e.PrivilegedMode = privilegedMode
}

func (e *Environment) SetEnvironmentType(environmentType string) {
    e.EnvironmentType = environmentType
}
