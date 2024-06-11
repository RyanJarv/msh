package awsapicallviacloudtrail

type Environment_1 struct {
    ComputeType string `json:"computeType,omitempty"`
    EnvironmentVariables []EnvironmentItem `json:"environmentVariables,omitempty"`
    Image string `json:"image,omitempty"`
    ImagePullCredentialsType string `json:"imagePullCredentialsType,omitempty"`
    PrivilegedMode bool `json:"privilegedMode,omitempty"`
    Environment_1Type string `json:"type,omitempty"`
}

func (e *Environment_1) SetComputeType(computeType string) {
    e.ComputeType = computeType
}

func (e *Environment_1) SetEnvironmentVariables(environmentVariables []EnvironmentItem) {
    e.EnvironmentVariables = environmentVariables
}

func (e *Environment_1) SetImage(image string) {
    e.Image = image
}

func (e *Environment_1) SetImagePullCredentialsType(imagePullCredentialsType string) {
    e.ImagePullCredentialsType = imagePullCredentialsType
}

func (e *Environment_1) SetPrivilegedMode(privilegedMode bool) {
    e.PrivilegedMode = privilegedMode
}

func (e *Environment_1) SetEnvironment1Type(environment1Type string) {
    e.Environment_1Type = environment1Type
}
