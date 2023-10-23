package fd

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io"
	"os"
)

// NewFd takes a file descriptor and resolves it to a remote queue if necessary.
//
// For example NewFd will may be called with os.Stdin in which it will read the first byte to determine
// if the first line is a json configuration which references a remote SQS queue. If the file descriptor is a TTY or the
// first character is a "{" we assume the first line contains a reference to a remote SQS queue that should be used for
// input.  In this case, an io.Reader is returned which reads from the SQS queue is used, otherwise we assume the input
// is data and return a bufio.Reader around the file descriptor that was passed.
//
// This allows us to pass references in the same way as normal data on stdin. Because we peek at the first line,
// the file descriptor passed here can not be used after this is called.
//
// TODO: To avoid ambiguity between configuration and data we use the shell to wrap output from all local commands in
//
//	an internally defined format.
func NewFd(file *os.File) (stdin interface{}, localStdin bool) {
	buf := bufio.NewReader(file)

	if utils.IsTTY(file) {
		return buf, true
	}

	return lo.Must2(resolveRef(buf))
}

func resolveRef(buf *bufio.Reader) (interface{}, bool, error) {
	peek, err := buf.Peek(1)
	if errors.Is(err, io.EOF) {
		L.Debug.Println("fd is closed, skipping config")
		return buf, true, nil
	} else if err != nil {
		return nil, true, fmt.Errorf("resolveRef: failed to peek: %w", err)
	}

	if peek[0] == '{' {
		fd, err := resolveFd(buf)
		if err != nil {
			return nil, true, err
		}

		return fd, false, err
	}

	return buf, true, nil
}

func resolveFd(buf *bufio.Reader) (io.Reader, error) {
	line, err := buf.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("resolveRef: failed to read line: %w", err)
	}

	var conf map[string]interface{}

	err = json.Unmarshal(line, &conf)
	if err != nil {
		return nil, fmt.Errorf("resolveRef: failed to unmarshal: %w", err)
	}

	url, ok := conf["SqsUrl"].(string)
	if !ok {
		L.Debug.Println("no SqsUrl found, using stdin")
		return buf, nil
	}

	L.Debug.Println("fd is sqs, skipping copy:", url)
	return OpenSqs(context.TODO(), url)
}
