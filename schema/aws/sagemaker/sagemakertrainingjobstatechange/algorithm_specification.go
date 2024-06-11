package sagemakertrainingjobstatechange

type AlgorithmSpecification struct {
    TrainingInputMode string `json:"TrainingInputMode,omitempty"`
    TrainingImage string `json:"TrainingImage,omitempty"`
}

func (a *AlgorithmSpecification) SetTrainingInputMode(trainingInputMode string) {
    a.TrainingInputMode = trainingInputMode
}

func (a *AlgorithmSpecification) SetTrainingImage(trainingImage string) {
    a.TrainingImage = trainingImage
}
