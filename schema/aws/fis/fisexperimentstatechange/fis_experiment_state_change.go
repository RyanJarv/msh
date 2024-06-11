package fisexperimentstatechange

type FISExperimentStateChange struct {
    NewState New_state `json:"new-state"`
    OldState Old_state `json:"old-state"`
    ExperimentId string `json:"experiment-id"`
}

func (f *FISExperimentStateChange) SetNewState(newState New_state) {
    f.NewState = newState
}

func (f *FISExperimentStateChange) SetOldState(oldState Old_state) {
    f.OldState = oldState
}

func (f *FISExperimentStateChange) SetExperimentId(experimentId string) {
    f.ExperimentId = experimentId
}
