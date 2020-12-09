package msh

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/pkg/errors"
	//"github.com/spf13/afero"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
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

func ReadConfig(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return StripShebang(string(file)), nil
}

// StripShebang strips the first line of input and returns the rest, an empty string will be returned if there is only
// one line.
func StripShebang(input string) (string) {
	all := strings.SplitN(input, "\n", 2)

	// Only splits the first newline
	if len(all) != 2 {
		return ""
	}

	return all[1]
}

func CleanName(path string) (s string) {
	s = filepath.Clean(path)
	s = strings.ReplaceAll(s, "/", "--")
	s = strings.ReplaceAll(s, "_", "-")
	s = regexp.MustCompile(`\.[^.]*$`).ReplaceAllString(s, "")
	return s
}

type TarFiles []struct {
	Name, Body string
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

func AwsIdentity(cfg Parser) *sts.GetCallerIdentityOutput {
	stsClient := sts.NewFromConfig(cfg.Aws.Config)
	identity, err := stsClient.GetCallerIdentity(cfg.Context, &sts.GetCallerIdentityInput{})
	if err != nil {
		panic(err)
	}
	return identity
}

func AwsIdentityTmpl(cfg Parser) sts.GetCallerIdentityOutput {
	awsIdentity := AwsIdentity(cfg)
	awsIdentity.ResultMetadata = middleware.Metadata{} // No need to make available to template
	return *awsIdentity
}


type Aws struct {
	Config   aws.Config `json:"-"`
	Identity sts.GetCallerIdentityOutput
}

type Global struct {
	Aws  Aws
	Args []string
	Project string
}

// TODO: For now we'll be passing this to templates, it contains creds... fix this before using pkg manager
type Parser struct {
	Path string
	Content string

	Args []string

	context.Context `json:"-"`
	Global
	Runnable 		`json:"-"`
}

type RunTemplate interface {
	Parsable
	Runnable
}

type Parsable interface {
	Parse(map[string]string)
}

type Runnable interface {
	Run() error
}
