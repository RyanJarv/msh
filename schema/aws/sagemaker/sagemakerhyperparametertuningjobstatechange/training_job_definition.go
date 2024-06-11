package sagemakerhyperparametertuningjobstatechange

type TrainingJobDefinition struct {
    AlgorithmSpecification AlgorithmSpecification `json:"AlgorithmSpecification"`
    StoppingCondition StoppingCondition `json:"StoppingCondition"`
    VpcConfig VpcConfig `json:"VpcConfig"`
    OutputDataConfig OutputDataConfig `json:"OutputDataConfig"`
    StaticHyperParameters Tags `json:"StaticHyperParameters"`
    ResourceConfig ResourceConfig `json:"ResourceConfig"`
    InputDataConfig []TrainingJobDefinitionItem `json:"InputDataConfig"`
    RoleArn string `json:"RoleArn"`
}

func (t *TrainingJobDefinition) SetAlgorithmSpecification(algorithmSpecification AlgorithmSpecification) {
    t.AlgorithmSpecification = algorithmSpecification
}

func (t *TrainingJobDefinition) SetStoppingCondition(stoppingCondition StoppingCondition) {
    t.StoppingCondition = stoppingCondition
}

func (t *TrainingJobDefinition) SetVpcConfig(vpcConfig VpcConfig) {
    t.VpcConfig = vpcConfig
}

func (t *TrainingJobDefinition) SetOutputDataConfig(outputDataConfig OutputDataConfig) {
    t.OutputDataConfig = outputDataConfig
}

func (t *TrainingJobDefinition) SetStaticHyperParameters(staticHyperParameters Tags) {
    t.StaticHyperParameters = staticHyperParameters
}

func (t *TrainingJobDefinition) SetResourceConfig(resourceConfig ResourceConfig) {
    t.ResourceConfig = resourceConfig
}

func (t *TrainingJobDefinition) SetInputDataConfig(inputDataConfig []TrainingJobDefinitionItem) {
    t.InputDataConfig = inputDataConfig
}

func (t *TrainingJobDefinition) SetRoleArn(roleArn string) {
    t.RoleArn = roleArn
}
