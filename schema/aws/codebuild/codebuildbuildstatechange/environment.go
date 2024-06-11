package codebuildbuildstatechange

type Environment struct {
    ComputeType string `json:"compute-type"`
    EnvironmentVariables []EnvironmentItem `json:"environment-variables"`
    Image string `json:"image"`
    ImagePullCredentialsType string `json:"image-pull-credentials-type"`
    PrivilegedMode bool `json:"privileged-mode"`
    EnvironmentType string `json:"type"`
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
