package utils

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
)

func IsNilOrEmpty(i interface{}) bool {
	switch v := i.(type) {
	case map[string]interface{}:
		return len(v) == 0
	default:
		return i == nil || reflect.ValueOf(i).IsNil()
	}
}

// JSONMarshal without HTML escaping
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

type EventAttributes struct {
	Cmd []string          `json:"cmd"`
	Env map[string]string `json:"env"`
}

func MapEnvVars(envVars []string) map[string]string {
	env := map[string]string{}
	for _, envVar := range envVars {
		parts := strings.SplitN(envVar, "=", 2)
		env[parts[0]] = parts[1]
	}
	return env
}
