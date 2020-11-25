package lib

import (
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

func ReadConfig(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	all := bytes.SplitN(file, []byte("\n"), 2)
	if len(all) != 2 {
		return "", errors.Wrap(err, "no content found")
	}
	return string(all[1]), err
}

func CleanName(path string) (s string) {
	s = filepath.Clean(path)
	s = strings.ReplaceAll(s, "/", "--")
	s = strings.ReplaceAll(s, "_", "-")
	s = regexp.MustCompile(`\.[^.]*$`).ReplaceAllString(s, "")
	return s
}

