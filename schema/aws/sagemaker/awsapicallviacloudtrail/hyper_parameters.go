package awsapicallviacloudtrail

type HyperParameters struct {
    EvalMetric string `json:"eval_metric,omitempty"`
    NumRound string `json:"num_round,omitempty"`
    Objective string `json:"objective,omitempty"`
}

func (h *HyperParameters) SetEvalMetric(evalMetric string) {
    h.EvalMetric = evalMetric
}

func (h *HyperParameters) SetNumRound(numRound string) {
    h.NumRound = numRound
}

func (h *HyperParameters) SetObjective(objective string) {
    h.Objective = objective
}
