package api

import (
	"embed"
	_ "embed"
)

//go:embed boto3/boto3/data
var boto3 embed.FS

type ResourceConfig struct {
	Service struct {
		Actions map[string]Action      `json:"actions"`
		HasMany map[string]Action      `json:"hasMany"`
		Has     map[string]interface{} `json:"has"`
	} `json:"service"`
	Resources map[string]Resource `json:"resources"`
}

type Resource struct {
	Identifiers []struct {
		Name       string `json:"name"`
		MemberName string `json:"memberName"`
	} `json:"identifiers"`
	Shape string `json:"shape"`
	Load  struct {
		Request struct {
			Operation string           `json:"operation"`
			Params    []IdentifierName `json:"params"`
		} `json:"request"`
		Path string `json:"path"`
	} `json:"load"`
	Actions      map[string]Action      `json:"actions"`
	BatchActions map[string]Action      `json:"batchActions"`
	Waiters      map[string]Waiter      `json:"waiters"`
	Has          map[string]SubResource `json:"has"`
	HasMany      struct {
		Volumes struct {
			Request  Request     `json:"request"`
			Resource SubResource `json:"resource"`
		} `json:"Volumes"`
		VpcAddresses struct {
			Request  Request     `json:"request"`
			Resource SubResource `json:"resource"`
		} `json:"VpcAddresses"`
	} `json:"hasMany"`
}

type Action struct {
	Request struct {
		Operation string `json:"operation"`
	} `json:"request"`
	Resource struct {
		Type        string           `json:"type"`
		Identifiers []IdentifierPath `json:"identifiers"`
		Path        string           `json:"path"`
	} `json:"resource"`
}

type Request struct {
	Operation string `json:"operation"`
	Params    []struct {
		Target string `json:"target"`
		Source string `json:"source"`
		Value  string `json:"value,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"params"`
}

type SubResource struct {
	Type        string           `json:"type"`
	Identifiers []IdentifierPath `json:"identifiers"`
}

type IdentifierPath struct {
	Target string `json:"target"`
	Source string `json:"source"`
	Path   string `json:"path"`
}

type IdentifierName struct {
	Target string `json:"target"`
	Source string `json:"source"`
	Name   string `json:"name"`
}

type Waiter struct {
	WaiterName string           `json:"waiterName"`
	Params     []IdentifierName `json:"params"`
	Path       string           `json:"path"`
}
