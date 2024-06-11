package awsapicallviacloudtrail

type ResponseElements struct {
    EndpointConfigArn string `json:"endpointConfigArn,omitempty"`
    TrainingJobArn string `json:"trainingJobArn,omitempty"`
    ModelArn string `json:"modelArn,omitempty"`
    TransformJobArn string `json:"transformJobArn,omitempty"`
}

func (r *ResponseElements) SetEndpointConfigArn(endpointConfigArn string) {
    r.EndpointConfigArn = endpointConfigArn
}

func (r *ResponseElements) SetTrainingJobArn(trainingJobArn string) {
    r.TrainingJobArn = trainingJobArn
}

func (r *ResponseElements) SetModelArn(modelArn string) {
    r.ModelArn = modelArn
}

func (r *ResponseElements) SetTransformJobArn(transformJobArn string) {
    r.TransformJobArn = transformJobArn
}
