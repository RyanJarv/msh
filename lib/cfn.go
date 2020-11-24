package lib

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	L "github.com/ryanjarv/msh/logger"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Cfn(argv []string) (err error) {
	path := filepath.Clean(argv[0])
	if err := CfnLint(path); err != nil {
		return err
	}

	var yaml string
	if yaml, err = ReadConfig(path); err != nil {
		return err
	}

	var client *cloudformation.Client
	if cfg, err := config.LoadDefaultConfig(); err != nil {
		return err
	} else {
		client = cloudformation.NewFromConfig(cfg)
	}

	ctx := context.Background()
	if role, err := client.CreateStack(ctx, &cloudformation.CreateStackInput{
		StackName:        aws.String(strings.Split(strings.ReplaceAll(path, "/", "-"), ".")[0]),
		Capabilities:     []types.Capability{types.CapabilityCapabilityIam},
		OnFailure:        types.OnFailureRollback,
		TemplateBody:     aws.String(yaml),
		TimeoutInMinutes: aws.Int32(1),
	}); err != nil {
		return err
	} else {
		L.Debug.Printf("Created role: %v", role)
	}
	return err
}

func CfnLint(path string) error {
	var cmd *exec.Cmd

	if abs, err := filepath.Abs("./config/cfn-lint"); err != nil {
		return err
	} else {
		cmd = exec.Command(abs, path)
	}
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	log.Printf("Running command and waiting for it to finish...")
	return cmd.Run()
}

