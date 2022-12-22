package fd

import (
	L "github.com/ryanjarv/msh/pkg/logger"
	"io"
	"os"
)

type HasEnv interface {
	Env() []string
}

type Referable interface {
	Arn() *string
}

func IsReferable(f interface{}) bool {
	_, ok := f.(Referable)
	return ok
}

// MustCopy Copy calls io.Copy and panics on error.
func MustCopy(dst io.WriteCloser, src io.Reader) {
	L.Debug.Printf("MustCopy: %+v %+v\n", dst, src)
	switch s := src.(type) {
	case *os.File:
		L.Debug.Printf("MustCopy: src (%T): %d %s\n", s, s.Fd(), s.Name())
	case Referable:
		L.Debug.Printf("MustCopy: src (%T): %s\n", s, *s.Arn())
	}

	switch d := dst.(type) {
	case *os.File:
		L.Debug.Printf("MustCopy: dst (%T): %d %s\n", d, d.Fd(), d.Name())
	case Referable:
		L.Debug.Printf("MustCopy: dst (%T): %s\n", d, *d.Arn())
	}

	w, err := io.Copy(dst, src)
	if err != nil {
		L.Error.Fatalln("MustCopy:", err)
	}

	err = dst.Close()
	if err != nil {
		L.Error.Fatalln("failed to close dst", err)
	}

	L.Debug.Printf("MustCopy: wrote %d bytes: err is %d\n", w, err)
}
