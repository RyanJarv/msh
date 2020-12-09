package msh

import (
	"log"
	"os"
	"os/exec"
)

// Deploy sets MSH_CONTEXT_TRIGGER_DEPLOY=true and executes the rest of the args. This typically will be a path to a msh lib config so
// this variable get's checked in Dockerfile.
func Deploy(argv []string) error {
	cmd := exec.Command(argv[0], argv[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Env = append(os.Environ(), "MSH_CONTEXT_TRIGGER_DEPLOY=true")
	log.Printf("Running command and waiting for it to finish...")
	return cmd.Run()
}
