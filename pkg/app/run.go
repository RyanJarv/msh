package app

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/logs"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"os"
	"os/exec"
	"strings"
)

func (s *App) Run() error {
	forceBuild := os.Getenv("MSH_BUILD") == "always"

	if utils.IsTTY(s.Stdout) || forceBuild {
		app := awscdk.NewApp(&awscdk.AppProps{})

		stack := awscdk.NewStack(app, aws.String(s.StackName()), &awscdk.StackProps{})

		if err := s.Compile(stack); err != nil {
			return err
		}

		return s.Build(app)
	} else {
		return s.State.WriteState(s.Stdout)
	}
}

func (s *App) Compile(scope constructs.Construct) error {
	L.Debug.Println("App.Compile")

	steps := BuildSteps(scope, s.State.Steps)

	for i, s := range steps {
		if v, ok := s.Step.(types.CdkStepCanInit); ok {
			logs.Debug("step %d: calling %s.Init(%s, %s)", i, s.Step, *s.Scope.ToString(), i)
			if err := v.Init(s.Scope, i); err != nil {
				return fmt.Errorf("%T.Init(): %w", s.Step, err)
			}
		}
	}

	for i, s := range steps {
		switch v := s.Step.(type) {
		case types.BeforeRunCdkStep:
			if err := v.BeforeRun(s); err != nil {
				logs.Debug("step %d: calling %T.Finalize(%s, %T, %T, %d)", i, s.Step, *s.Scope.ToString(), s.Last, s.Next, i)
				return fmt.Errorf("%T.Finalize(...): %w", s.Step, err)
			}
		}
	}

	var lastFoundRole awsiam.IRole
	for i, s := range steps {
		switch v := s.Step.(type) {
		case types.RunnableCdkStep:
			if err := v.Run(s); err != nil {
				logs.Debug("step %d: calling %T.Run(%s, %T, %T, %d)", i, s.Step, *s.Scope.ToString(), s.Last, s.Next, i)
				return fmt.Errorf("%T.Run(...): %w", s.Step, err)
			}
		}

		if v, ok := s.Step.(types.AcceptsLastRole); ok {
			logs.Debug("step %d: calling %s.LastFoundRole(%s)", i, s.Step, lastFoundRole)
			v.LastFoundRole(scope, lastFoundRole)
		}

		if v, ok := s.Step.(types.HasRole); ok {
			logs.Debug("step %d: calling %s.Role()", i, s.Step)
			lastFoundRole = v.Role()
		}
	}

	for i, s := range steps {
		switch v := s.Step.(type) {
		case types.AfterRunCdkStep:
			if err := v.AfterRun(s); err != nil {
				logs.Debug("step %d: calling %T.Finalize(%s, %T, %T, %d)", i, s.Step, *s.Scope.ToString(), s.Last, s.Next, i)
				return fmt.Errorf("%T.Finalize(...): %w", s.Step, err)
			}
		}
	}

	return nil
}

func BuildSteps(scope constructs.Construct, input []types.Step) []*types.StepRunInfo {
	steps := []*types.StepRunInfo{}

	for i, v := range input {
		this := &types.StepRunInfo{
			Index: i,
			Id: func(s ...string) *string {
				return aws.String(strings.Join(s, "-") + *utils.GetId(v.GetName(), i))
			},
			Step:  v.Value.(types.CdkStep),
			Scope: constructs.NewConstruct(scope, jsii.String(fmt.Sprintf("%T%d", v, i))),
			Last:  nil,
			Next:  nil,
		}

		if i > 0 {
			steps[i-1].Next = this
		}
		if i < len(steps) {
			steps[i+1].Last = this

		}
		steps = append(steps, this)
	}
	return steps
}

func (s *App) Build(app awscdk.App) error {
	L.Debug.Println("App.Build")
	synth := app.Synth(&awscdk.StageSynthesisOptions{})
	if synth == nil || synth.Stacks() == nil || len(*synth.Stacks()) != 1 {
		return fmt.Errorf("build: failed to synthesize app: %v", synth)
	}

	for _, stack := range *synth.Stacks() {
		tmpl := string(lo.Must(json.Marshal(stack.Template())))
		L.Debug.Println(tmpl)
		if os.Getenv("MSH_TEMPLATE") != "" {
			fmt.Println(tmpl)
		}
	}

	L.Debug.Println("synth directory:", *synth.Directory())

	if skip := os.Getenv("MSH_SKIPDEPLOY") != ""; !skip {
		if err := Deploy(synth); err != nil {
			return fmt.Errorf("build: failed to deploy: %w", err)
		}
	}

	return nil
}

func Deploy(synth cxapi.CloudAssembly) error {
	args := []string{"bootstrap", "--app", *synth.Directory()}
	L.Debug.Println("Running: cdk", strings.Join(args, " "))

	cmd := exec.Command("cdk", args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build: failed to bootstrap: %w", err)
	}

	cmd = exec.Command(
		"cdk", "deploy",
		"--hotswap-fallback",
		"--require-approval=never",
		"--app", *synth.Directory(),
	)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build: failed to bootstrap: %w", err)
	}

	return nil
}
