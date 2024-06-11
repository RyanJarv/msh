package sagemakergroundtruthlabelingjobtaskstatuschange

type SageMakerGroundTruthLabelingJobTaskStatusChange struct {
    State string `json:"State"`
    LabelingJobArn string `json:"labeling-job-arn"`
    WorkteamTeamArn string `json:"workteam-team-arn"`
}

func (s *SageMakerGroundTruthLabelingJobTaskStatusChange) SetState(state string) {
    s.State = state
}

func (s *SageMakerGroundTruthLabelingJobTaskStatusChange) SetLabelingJobArn(labelingJobArn string) {
    s.LabelingJobArn = labelingJobArn
}

func (s *SageMakerGroundTruthLabelingJobTaskStatusChange) SetWorkteamTeamArn(workteamTeamArn string) {
    s.WorkteamTeamArn = workteamTeamArn
}
