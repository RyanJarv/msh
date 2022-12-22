package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	ProgName = "[" + os.Args[0] + "]"
	Debug    = log.New(os.Stderr, ProgName+" [DEBUG] ", log.Flags())
	Info     = log.New(os.Stderr, ProgName+" [INFO] ", log.Flags())
	Error    = log.New(os.Stderr, ProgName+" [ERROR] ", log.Flags())
)

func init() {
	if os.Getenv("DEBUG") == "" {
		if f, err := os.Open(os.DevNull); err != nil {
			fmt.Errorf("could not open %s, debug output will print to stderr", os.DevNull)
		} else {
			Debug.SetOutput(f)
		}
	}
}
