package utils

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/samber/lo"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Bool bool

func (b *Bool) Capture(v []string) error { *b = v[0] == "true"; return nil }

type Value struct {
	Boolean *Bool    `@("true"|"false")`
	String  *string  `| @(String|Char|RawString|Ident)`
	Number  *float64 `| @(Float|Int)`
	Array   []*Value `| "[" ( @@ ","? )* "]"`
}

func (l *Value) ToString() string {
	switch {
	case l.Boolean != nil:
		return fmt.Sprintf("%t", *l.Boolean)
	case l.String != nil:
		return *l.String
	case l.Number != nil:
		return fmt.Sprintf("%d", *l.Number)
	case l.Array != nil:
		return strings.Join(lo.Map(l.Array, func(v *Value, _ int) string {
			return v.ToString()
		}), " ")
	}
	panic("??")
}

func (l *Value) GoString() string {
	switch {
	case l.Boolean != nil:
		return fmt.Sprintf("%v", *l.Boolean)
	case l.String != nil:
		return fmt.Sprintf("%q", *l.String)
	case l.Number != nil:
		return fmt.Sprintf("%v", *l.Number)
	case l.Array != nil:
		out := []string{}
		for _, v := range l.Array {
			out = append(out, v.GoString())
		}
		return fmt.Sprintf("[]*Value{ %s }", strings.Join(out, ", "))
	}
	panic("??")
}

type Step struct {
	Key    string   `@Ident`
	Params []*Value `@@*`
	Block  []*Step  `("{" (@@ ("|"|";")?)* "}")?`
}

func (s *Step) GetArgs() []string {
	return append([]string{s.Key}, lo.Map(s.Params, func(v *Value, _ int) string {
		return v.ToString()
	})...)
}

func (s *Step) ParamString() string {
	return strings.Join(lo.Map(s.Params, func(v *Value, _ int) string {
		return v.ToString()
	}), ".")
}

func (s *Step) Name() string {
	return s.Key + "." + s.ParamString()
}

type Msh struct {
	Pipelines []*Step `@@*`
}

var parser = participle.MustBuild[Msh](participle.Unquote())

func ParseFile(path string) (*Msh, error) {
	var script string
	if path == "" {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("read stdin: %w", err)
		}

		script = string(b)
		path = "(stdin)"
	} else {
		file, err := os.ReadFile(filepath.Clean(path))
		if err != nil {
			return nil, fmt.Errorf("read file: %w", err)
		}

		script = string(file)
	}

	return parser.ParseString(path, script)
}
