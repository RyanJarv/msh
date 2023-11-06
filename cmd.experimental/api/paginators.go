package api

import (
	"embed"
	_ "embed"
)

//go:embed botocore/botocore/data
var botocore embed.FS

type PaginatorConfig struct {
	Pagination map[string]Paginator `json:"pagination"`
}

type Paginator struct {
	InputToken  string      `json:"input_token"`
	OutputToken string      `json:"output_token"`
	LimitKey    string      `json:"limit_key"`
	ResultKey   interface{} `json:"result_key"` // Can be a list or a string.
}

//var BrokenPaginators = map[string]string{
//	"DescribeNetworkInterfaces": "$.outputDetails.truncated",
//}
