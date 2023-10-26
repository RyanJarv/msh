package api

import (
	"encoding/json"
	"testing"
)

func TestPaginatorConfig(t *testing.T) {
	file, err := botocore.ReadFile("botocore/botocore/data/ec2/2016-11-15/paginators-1.json")
	if err != nil {
		t.Fatal(err)
	}

	config := PaginatorConfig{}
	if err := json.Unmarshal(file, &config); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}

	if got := config.Pagination["DescribeInstances"].InputToken; got != "NextToken" {
		t.Errorf("DescribeInstances result key: got %s, want Reservations", got)
	}
}
