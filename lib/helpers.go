package lib

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func GetRandStr() string {
	rand.Seed(time.Now().UnixNano())
	tag := strconv.Itoa(rand.Int())
	return tag
}

func MakeTarGz(files *TarFiles) io.Reader {
	// Create and add some files to the archive.
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	for _, file := range *files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
			Typeflag: tar.TypeReg,
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		panic(err)
	}

	gzipOut := bytes.NewBuffer([]byte{})
	writer := gzip.NewWriter(gzipOut)
	defer writer.Close()

	_, err := buf.WriteTo(writer)
	if err != nil {
		panic(err)
	}

	return gzipOut
}

