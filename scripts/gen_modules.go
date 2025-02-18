package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type Module struct {
	Path string
	Name string
}

type Modules struct {
	Modules []Module
}

var templ = `
package main

import (
	"flag"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/app"
	"strings"

	{{- range .Modules}}
	"github.com/ryanjarv/msh/{{.Path}}"
	{{- end}}
)

var Registry = types.Registry{
	{{- range .Modules}}
	"{{.Name}}": {{.Name}}.New,
	{{- end}}
}

func GetStepFunc(cmdName string) (func (*app.App, []string) (types.CdkStep, error)) {
    flag.Parse()

	cmdName = strings.Trim(cmdName, "@.")

	switch cmdName {
	{{- range .Modules}}
	case "{{.Name}}":
		 return {{.Name}}.New
	{{- end}}
	default:
		return nil
	}
}
`

func main() {
	m := Modules{}

	for _, p := range []string{"../cmd", "../cmd.experimental"} {
		dir, err := os.ReadDir(p)
		if err != nil {
			log.Fatalf("read dir: %s", err)
		}
		for _, f := range dir {
			if f.IsDir() {
				m.Modules = append(m.Modules, Module{
					Path: filepath.Join(filepath.Base(p), f.Name()),
					Name: f.Name(),
				})
			}
		}
	}

	t, err := template.New("modules").Parse(templ)
	if err != nil {
		log.Fatalf("template: %s", err)
	}

	f, err := os.Create("modules.go")
	if err != nil {
		log.Fatalf("open: %s", err)
	}

	if err = t.Execute(f, m); err != nil {
		log.Fatalf("execute: %s", err)
	}
}
