package utils

import (
	"fmt"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"strings"
)

var DevNull, _ = os.Open(os.DevNull)

func IsTTY(file *os.File) bool {
	if os.Getenv("MSH_TTY") == "false" {
		return false
	}

	envName := fmt.Sprintf("MSH_TTY_%d", int(file.Fd()))
	if env := os.Getenv(envName); env == "false" {
		return false
	} else if env == "true" {
		return true
	}

	stat, err := file.Stat()
	if err != nil {
		L.Debug.Printf("fd %d: stat failed, assuming we're not using a tty", file.Fd())
		return false
	}

	if stat.Mode()&os.ModeCharDevice == 0 {
		L.Debug.Printf("fd %d: pipe\n", file.Fd())
		return false
	} else {
		L.Debug.Printf("fd %d: tty\n", file.Fd())
		return true
	}
}

func ExpandPath(s string) string {
	if strings.HasPrefix(s, "~") {
		return strings.Replace(s, "~", lo.Must(os.UserHomeDir()), -1)
	}
	return lo.Must(filepath.Abs(s))
}
