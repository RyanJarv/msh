package awsapicallviacloudtrail

type RequestParameters struct {
    PrimaryContainer PrimaryContainer `json:"primaryContainer,omitempty"`
    StoppingCondition StoppingCondition `json:"stoppingCondition,omitempty"`
    OutputDataConfig OutputDataConfig `json:"outputDataConfig,omitempty"`
    TransformInput TransformInput `json:"transformInput,omitempty"`
    AlgorithmSpecification AlgorithmSpecification `json:"algorithmSpecification,omitempty"`
    ResourceConfig ResourceConfig `json:"resourceConfig,omitempty"`
    TransformOutput TransformOutput `json:"transformOutput,omitempty"`
    HyperParameters HyperParameters `json:"hyperParameters,omitempty"`
    TransformResources TransformResources `json:"transformResources,omitempty"`
    EndpointConfigName string `json:"endpointConfigName,omitempty"`
    ProductionVariants []RequestParametersItem `json:"productionVariants,omitempty"`
    EndpointName string `json:"endpointName,omitempty"`
    EnableNetworkIsolation bool `json:"enableNetworkIsolation,omitempty"`
    Tags []interface{} `json:"tags,omitempty"`
    ModelName string `json:"modelName,omitempty"`
    ExecutionRoleArn string `json:"executionRoleArn,omitempty"`
    TrainingJobName string `json:"trainingJobName,omitempty"`
    RoleArn string `json:"roleArn,omitempty"`
    EnableManagedSpotTraining bool `json:"enableManagedSpotTraining,omitempty"`
    InputDataConfig []RequestParametersItem_2 `json:"inputDataConfig,omitempty"`
    EnableInterContainerTrafficEncryption bool `json:"enableInterContainerTrafficEncryption,omitempty"`
    TransformJobName string `json:"transformJobName,omitempty"`
}

func (r *RequestParameters) SetPrimaryContainer(primaryContainer PrimaryContainer) {
    r.PrimaryContainer = primaryContainer
}

func (r *RequestParameters) SetStoppingCondition(stoppingCondition StoppingCondition) {
    r.StoppingCondition = stoppingCondition
}

func (r *RequestParameters) SetOutputDataConfig(outputDataConfig OutputDataConfig) {
    r.OutputDataConfig = outputDataConfig
}

func (r *RequestParameters) SetTransformInput(transformInput TransformInput) {
    r.TransformInput = transformInput
}

func (r *RequestParameters) SetAlgorithmSpecification(algorithmSpecification AlgorithmSpecification) {
    r.AlgorithmSpecification = algorithmSpecification
}

func (r *RequestParameters) SetResourceConfig(resourceConfig ResourceConfig) {
    r.ResourceConfig = resourceConfig
}

func (r *RequestParameters) SetTransformOutput(transformOutput TransformOutput) {
    r.TransformOutput = transformOutput
}

func (r *RequestParameters) SetHyperParameters(hyperParameters HyperParameters) {
    r.HyperParameters = hyperParameters
}

func (r *RequestParameters) SetTransformResources(transformResources TransformResources) {
    r.TransformResources = transformResources
}

func (r *RequestParameters) SetEndpointConfigName(endpointConfigName string) {
    r.EndpointConfigName = endpointConfigName
}

func (r *RequestParameters) SetProductionVariants(productionVariants []RequestParametersItem) {
    r.ProductionVariants = productionVariants
}

func (r *RequestParameters) SetEndpointName(endpointName string) {
    r.EndpointName = endpointName
}

func (r *RequestParameters) SetEnableNetworkIsolation(enableNetworkIsolation bool) {
    r.EnableNetworkIsolation = enableNetworkIsolation
}

func (r *RequestParameters) SetTags(tags []interface{}) {
    r.Tags = tags
}

func (r *RequestParameters) SetModelName(modelName string) {
    r.ModelName = modelName
}

func (r *RequestParameters) SetExecutionRoleArn(executionRoleArn string) {
    r.ExecutionRoleArn = executionRoleArn
}

func (r *RequestParameters) SetTrainingJobName(trainingJobName string) {
    r.TrainingJobName = trainingJobName
}

func (r *RequestParameters) SetRoleArn(roleArn string) {
    r.RoleArn = roleArn
}

func (r *RequestParameters) SetEnableManagedSpotTraining(enableManagedSpotTraining bool) {
    r.EnableManagedSpotTraining = enableManagedSpotTraining
}

func (r *RequestParameters) SetInputDataConfig(inputDataConfig []RequestParametersItem_2) {
    r.InputDataConfig = inputDataConfig
}

func (r *RequestParameters) SetEnableInterContainerTrafficEncryption(enableInterContainerTrafficEncryption bool) {
    r.EnableInterContainerTrafficEncryption = enableInterContainerTrafficEncryption
}

func (r *RequestParameters) SetTransformJobName(transformJobName string) {
    r.TransformJobName = transformJobName
}
