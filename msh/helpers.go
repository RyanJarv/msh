package msh

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

// InSliceStr will return the index of str in slice if it exists, otherwise it will return nil
func InSliceStr(slice []string, str string) *int {
	for i, val := range slice {
		if val == str {
			return &i
		}
	}
	return nil
}

type StackTracer interface {
	StackTrace() errors.StackTrace
}

func GetMshContext() *MshContext {
	var context MshContext
	if v := os.Getenv(MSH_CONTEXT); v != "" {
		if err := json.Unmarshal([]byte(v), &context); err != nil {
			panic(err)
		}
		return &context
	}
	return nil
}
