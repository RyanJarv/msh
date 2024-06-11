package sagemakertransformjobstatechange

type SageMakerTransformJobStateChange struct {
    TransformResources TransformResources `json:"TransformResources"`
    TransformOutput TransformOutput `json:"TransformOutput"`
    TransformInput TransformInput `json:"TransformInput"`
    Tags []Tags `json:"Tags,omitempty"`
    TransformStartTime float64 `json:"TransformStartTime,omitempty"`
    ModelName string `json:"ModelName"`
    CreationTime float64 `json:"CreationTime"`
    TransformJobArn string `json:"TransformJobArn"`
    TransformJobStatus string `json:"TransformJobStatus"`
    TransformJobName string `json:"TransformJobName"`
    TransformEndTime float64 `json:"TransformEndTime,omitempty"`
}

func (s *SageMakerTransformJobStateChange) SetTransformResources(transformResources TransformResources) {
    s.TransformResources = transformResources
}

func (s *SageMakerTransformJobStateChange) SetTransformOutput(transformOutput TransformOutput) {
    s.TransformOutput = transformOutput
}

func (s *SageMakerTransformJobStateChange) SetTransformInput(transformInput TransformInput) {
    s.TransformInput = transformInput
}

func (s *SageMakerTransformJobStateChange) SetTags(tags []Tags) {
    s.Tags = tags
}

func (s *SageMakerTransformJobStateChange) SetTransformStartTime(transformStartTime float64) {
    s.TransformStartTime = transformStartTime
}

func (s *SageMakerTransformJobStateChange) SetModelName(modelName string) {
    s.ModelName = modelName
}

func (s *SageMakerTransformJobStateChange) SetCreationTime(creationTime float64) {
    s.CreationTime = creationTime
}

func (s *SageMakerTransformJobStateChange) SetTransformJobArn(transformJobArn string) {
    s.TransformJobArn = transformJobArn
}

func (s *SageMakerTransformJobStateChange) SetTransformJobStatus(transformJobStatus string) {
    s.TransformJobStatus = transformJobStatus
}

func (s *SageMakerTransformJobStateChange) SetTransformJobName(transformJobName string) {
    s.TransformJobName = transformJobName
}

func (s *SageMakerTransformJobStateChange) SetTransformEndTime(transformEndTime float64) {
    s.TransformEndTime = transformEndTime
}
