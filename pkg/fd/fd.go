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

type Fd struct {
	io.WriterTo
	io.Closer
	file *os.File
}

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

	url, ok := conf["url"].(string)
	if !ok {
		L.Debug.Println("no url found, using stdin")
		return buf, nil
	}

	L.Debug.Println("fd is sqs, skipping copy:", url)
	return OpenSqs(context.TODO(), url)
}

func readLine(f io.Reader) ([]byte, error) {
	b := make([]byte, 1, 1)
	var s []byte
	for {
		_, err := f.Read(b)
		if err != nil {
			return nil, fmt.Errorf("read byte: %w", err)
		}
		if b[0] == '\n' {
			break
		}

		s = append(s, b...)

	}

	return s, nil
}
