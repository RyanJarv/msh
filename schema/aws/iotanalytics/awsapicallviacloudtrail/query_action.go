package awsapicallviacloudtrail

type QueryAction struct {
    SqlQuery string `json:"sqlQuery,omitempty"`
}

func (q *QueryAction) SetSqlQuery(sqlQuery string) {
    q.SqlQuery = sqlQuery
}
