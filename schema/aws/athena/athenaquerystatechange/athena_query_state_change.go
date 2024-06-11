package athenaquerystatechange

type AthenaQueryStateChange struct {
    CurrentState string `json:"currentState"`
    PreviousState string `json:"previousState"`
    QueryExecutionId string `json:"queryExecutionId"`
    SequenceNumber string `json:"sequenceNumber"`
    StatementType string `json:"statementType"`
    VersionId string `json:"versionId"`
    WorkgroupName string `json:"workgroupName"`
}

func (a *AthenaQueryStateChange) SetCurrentState(currentState string) {
    a.CurrentState = currentState
}

func (a *AthenaQueryStateChange) SetPreviousState(previousState string) {
    a.PreviousState = previousState
}

func (a *AthenaQueryStateChange) SetQueryExecutionId(queryExecutionId string) {
    a.QueryExecutionId = queryExecutionId
}

func (a *AthenaQueryStateChange) SetSequenceNumber(sequenceNumber string) {
    a.SequenceNumber = sequenceNumber
}

func (a *AthenaQueryStateChange) SetStatementType(statementType string) {
    a.StatementType = statementType
}

func (a *AthenaQueryStateChange) SetVersionId(versionId string) {
    a.VersionId = versionId
}

func (a *AthenaQueryStateChange) SetWorkgroupName(workgroupName string) {
    a.WorkgroupName = workgroupName
}
