package sagemakerhyperparametertuningjobstatechange

type AlgorithmSpecification struct {
    TrainingInputMode string `json:"TrainingInputMode"`
    MetricDefinitions []AlgorithmSpecificationItem `json:"MetricDefinitions"`
    TrainingImage string `json:"TrainingImage"`
}

func (a *AlgorithmSpecification) SetTrainingInputMode(trainingInputMode string) {
    a.TrainingInputMode = trainingInputMode
}

func (a *AlgorithmSpecification) SetMetricDefinitions(metricDefinitions []AlgorithmSpecificationItem) {
    a.MetricDefinitions = metricDefinitions
}

func (a *AlgorithmSpecification) SetTrainingImage(trainingImage string) {
    a.TrainingImage = trainingImage
}
