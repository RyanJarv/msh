package redshiftdatastatementstatuschangeevent

import (
    "time"
)


type RedshiftDataStatementStatusChange struct {
    ExpireAt time.Time `json:"expireAt"`
    Principal string `json:"principal"`
    RedshiftQueryId float64 `json:"redshiftQueryId"`
    Rows float64 `json:"rows"`
    State string `json:"state"`
    StatementId string `json:"statementId"`
    StatementName string `json:"statementName"`
}

func (r *RedshiftDataStatementStatusChange) SetExpireAt(expireAt time.Time) {
    r.ExpireAt = expireAt
}

func (r *RedshiftDataStatementStatusChange) SetPrincipal(principal string) {
    r.Principal = principal
}

func (r *RedshiftDataStatementStatusChange) SetRedshiftQueryId(redshiftQueryId float64) {
    r.RedshiftQueryId = redshiftQueryId
}

func (r *RedshiftDataStatementStatusChange) SetRows(rows float64) {
    r.Rows = rows
}

func (r *RedshiftDataStatementStatusChange) SetState(state string) {
    r.State = state
}

func (r *RedshiftDataStatementStatusChange) SetStatementId(statementId string) {
    r.StatementId = statementId
}

func (r *RedshiftDataStatementStatusChange) SetStatementName(statementName string) {
    r.StatementName = statementName
}
