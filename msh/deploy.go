package msh

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	L "github.com/ryanjarv/msh/logger"
	"os"
	"os/exec"
)

// Remote sets MSH_CONTEXT_TRIGGER_DEPLOY=true and executes the rest of the args. This typically will be a path to a msh lib config so
// this variable get's checked in Dockerfile.
func Remote(argv []string) {
	msg := fmt.Sprintf("executing cmd %s with args %v", argv[1], argv[2:])
	L.Debug.Println(msg)

	// Re-execute removing remote arg
	cmd := exec.Command(argv[1], argv[2:]...)

	// This doesn't work for some reason
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

	// Store remote arg in MSH_CONTEXT environment
	json, err := json.Marshal(MshContext{
		RunEnv: RunEnv(argv[0]),
	})
	if err != nil {
		panic(err)
	}

	L.Debug.Printf("SettingEnv: MSH_CONTEXT='%s'", string(json))

	cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", MSH_CONTEXT, string(json)))

	L.Debug.Printf("Running command and waiting for it to finish...")
	if err := cmd.Run(); err != nil {
		panic(errors.New("failed " + msg))
	}
}
