package utils

import (
	"bytes"
	L "github.com/ryanjarv/msh/logger"
	"text/template"
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

func Template(params interface{}, templates ...string) ([]string, error) {
	L.Debug.Printf("un-formatted templates: %v", templates)
	var result []string

	for _, t := range templates {
		tmpl, err := template.New("").Parse(t)
		if err != nil {
			return result, Wrap(err, "Template")
		}

		w := bytes.NewBuffer([]byte{})
		err = tmpl.Execute(w, params)
		if err != nil {
			return result, Wrap(err, "Template")
		}

		result = append(result, w.String())
	}

	L.Debug.Printf("formatted templates: %v", result)
	return result, nil
}

func Append[T comparable](s ...T) []T {
	return s
}
