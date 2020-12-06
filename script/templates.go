package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const TemplateDir = "./templates"
const OutputFile = "gen/templates.go"

// Reads all .txt files in the current folder
// and encodes them as strings literals in textfiles.go
func main() {
	fs, err := ioutil.ReadDir(TemplateDir)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(OutputFile)
	if err != nil {
		panic(err)
	}

	WriteTo(out, "package gen \n\nconst (\n")

	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".tmpl") {
			name := strings.TrimSuffix(f.Name(), ".tmpl")
			WriteTo(out, strings.ToTitle(name) + " = `")

			f, err := os.Open(filepath.Join(TemplateDir, f.Name()))
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(out, f)
			if err != nil {
				panic(err)
			}

			WriteTo(out,"`)\n")
		}
	}
}

func WriteTo(out *os.File, s string) {
	if _, err := out.Write([]byte(s)); err != nil {
		panic(err)
	}
}
