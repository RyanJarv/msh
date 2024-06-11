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
	"fmt"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/utils"

	{{- range .Modules}}
	"github.com/ryanjarv/msh/{{.Path}}"
	{{- end}}
)

var Registry = types.NewRegistry(
	{{- range .Modules}}
	{{.Name}}.New,
	{{- end}}
)

func NewModule(app app.App, name string) (step types.CdkStep, err error) {
    flag.Parse()

	switch name {
	{{- range .Modules}}
	case ".{{.Name}}":
		 step, err = {{.Name}}.New(app)
	{{- end}}
	default:
		err = fmt.Errorf("unknown module: %s", name)
	}

	return step, utils.Wrap(err, name)
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
