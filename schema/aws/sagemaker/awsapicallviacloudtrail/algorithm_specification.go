package awsapicallviacloudtrail

type AlgorithmSpecification struct {
    TrainingInputMode string `json:"trainingInputMode,omitempty"`
    TrainingImage string `json:"trainingImage,omitempty"`
}

func (a *AlgorithmSpecification) SetTrainingInputMode(trainingInputMode string) {
    a.TrainingInputMode = trainingInputMode
}

func (a *AlgorithmSpecification) SetTrainingImage(trainingImage string) {
    a.TrainingImage = trainingImage
}
