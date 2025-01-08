package utils

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/samber/lo"
	"reflect"
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

type EventAttributes struct {
	Cmd []string          `json:"cmd"`
	Env map[string]string `json:"env"`
}

func UnTitle(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}

func GetId(name string, i int) *string {
	return aws.String(fmt.Sprintf("%s-%d", name, i))
}
