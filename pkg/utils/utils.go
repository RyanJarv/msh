package utils

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/samber/lo"
	"reflect"
	"strings"
	"unicode"
)

func ParseArgs(args []string) *flag.FlagSet {
	flagset := flag.NewFlagSet("args", flag.ExitOnError)
	lo.Must0(flagset.Parse(args))
	return flagset
}

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

func EachAs[T any](values []any) ([]T, bool) {
	var out []T
	for _, v := range values {
		o, ok := v.(T)
		if !ok {
			return out, false
		}

		out = append(out, o)
	}

	return out, true
}

func Empty[T any](t T) T {
	var zero T
	return zero
}

func UnTitle(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}
