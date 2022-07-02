package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var DevNull, _ = os.Open(os.DevNull)

func NewTmpDir(id string) (TmpDir, error) {
	tmpDir, err := ioutil.TempDir(os.TempDir(), "pkg-"+id)
	return TmpDir{Path: tmpDir}, Wrap(err, "NewTmpDir: %w", err)
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

func (t TmpDir) Remove() error {
	err := os.RemoveAll(t.Path)
	return Wrap(err, "pagh %s", t.Path)
}
