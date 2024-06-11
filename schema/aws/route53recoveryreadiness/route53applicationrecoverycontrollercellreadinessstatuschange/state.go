package route53applicationrecoverycontrollercellreadinessstatuschange

type ReadinessStatus string

// Enum of readinessStatus
const (
    READY ReadinessStatus = "READY"
    NOT_READY ReadinessStatus = "NOT_READY"
    UNKNOWN ReadinessStatus = "UNKNOWN"
    NOT_AUTHORIZED ReadinessStatus = "NOT_AUTHORIZED"
)

type State struct {
    ReadinessStatus ReadinessStatus `json:"readiness-status"`
}

func (s *State) SetReadinessStatus(readinessStatus ReadinessStatus) {
    s.ReadinessStatus = readinessStatus
}
