package msh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	L "github.com/ryanjarv/msh/logger"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

// TODO: Possibly remove
func DockerTmplExecNoContext(path string, args []string) {
	L.Debug.Printf("Running DockerTmplExecNoContext with path '%s' and args '%v'\n", path, args)

	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	params := ExecParams{
		Path: absPath,
		Name: filepath.Base(absPath),
	}

	dir := NewTmpDir("docker")
	defer dir.Remove()
	params.ContextDir = dir.Path

	TmplExec(params, args)
}

func NewTmpDir(id string) TmpDir {
	tmpDir, err := ioutil.TempDir(os.TempDir(), "msh-" + id)
	if err != nil {
		panic(err)
	}
	return TmpDir{
		Path: tmpDir,
	}
}

type TmpDir struct {
	Path string
}

func (t TmpDir) AddFile(src, dst string) {
	if err := os.Link(src, filepath.Join(t.Path, dst)); err != nil {
		panic(err)
	}
}

func (t TmpDir) AddContent(content []byte, dst string) {
	if err := ioutil.WriteFile(filepath.Join(t.Path, dst), content, os.FileMode(0400)); err != nil {
		panic(err)
	}
}

func (t TmpDir) Remove() {
	if err := os.RemoveAll(t.Path); err != nil {
		panic(err)
	}
}

func TmplExec(params interface{}, args []string) {
	var fmtArgs []string
	for _, arg := range args {
		fmtArgs = append(fmtArgs, Tmpl(params, arg))
	}

	L.Debug.Printf("Running: %v", fmtArgs)

	cmd := exec.Command(fmtArgs[0], fmtArgs[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Env = append(cmd.Env, os.Environ()...)

	// We look for the app arg separator and start adding shell style environment
	// variables to make available to the compose file.
	if result := InSliceStr(args, "--"); result != nil {
		for i, arg := range args[*result:] {
			cmd.Env = append(cmd.Env, fmt.Sprintf("ARG_%d=%s", i, arg))
		}
	}

	log.Printf("Running command and waiting for it to finish...")
	if err := cmd.Run(); err != nil {
		cmd.StderrPipe()
		panic(err)
	}
}

func Tmpl(params interface{}, s string) string {
	tmpl, err := template.New("").Parse(s)
	if err != nil {
		panic(err)
	}

	w := bytes.NewBuffer([]byte{})
	if err := tmpl.Execute(w, params); err != nil {
		panic(err)
	}

	return w.String()
}
