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
	fdDebug(src, "dst")
	fdDebug(dst, "dst")

	w, err := io.Copy(dst, src)
	if err != nil {
		L.Error.Fatalln("MustCopy:", err)
	}

	err = dst.Close()
	if err != nil {
		L.Error.Fatalln("closing dst", err)
	}

	L.Debug.Printf("MustCopy: wrote %d bytes: err is %d\n", w, err)
}

func fdDebug(fd interface{}, msg string) {
	switch d := fd.(type) {
	case *os.File:
		L.Debug.Printf("MustCopy: %s (%T): %d %s\n", msg, d, d.Fd(), d.Name())
	case Referable:
		L.Debug.Printf("MustCopy: %s (%T): %s\n", msg, d, *d.Arn())
	default:
		L.Debug.Printf("MustCopy: %s (%T)\n", msg, d)
	}
}
