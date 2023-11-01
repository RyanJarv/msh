package app

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io"
	"os"
	"os/exec"
)

func (a *App) Run(step types.IStep) error {
	a.State.AddStep(step)

	if !utils.IsTTY(a.Stdout) {
		return a.State.WriteState()
	}

	if utils.IsTTY(a.Stdout) || os.Getenv("MSH_BUILD") == "always" {
		app := awscdk.NewApp(&awscdk.AppProps{})
		stack := awscdk.NewStack(app, jsii.String("msh"), &awscdk.StackProps{})

		_, err := a.Compile(stack, nil, 0)
		if err != nil {
			return fmt.Errorf("compile: %w", err)
		}

		return a.Build(app)
	}

	return nil
}

func (s *State) WriteState() error {
	d, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("writeState: failed to marshal app: %w", err)
	}

	_, err = io.WriteString(os.Stdout, string(d)+"\n")
	if err != nil {
		return fmt.Errorf("writeState: failed to write app: %w", err)
	}

	return nil
}

func (a *App) Compile(stack constructs.Construct, next []interface{}, i int) ([]interface{}, error) {
	// next represents the output of the last Step, which will be passed to the next.
	if next == nil {
		next = []interface{}{}
	}

	// Reverse the steps so the source receives the next Step instead of the previous one.
	steps := lo.Reverse(a.State.Steps)

	for i, step := range steps {
		L.Debug.Println("running Step:", step.GetName())

		this := constructs.NewConstruct(stack, jsii.String(fmt.Sprintf("%s-%d", step.GetName(), i)))

		switch s := step.Value.(type) {
		case types.CdkStep:
			var n interface{}
			if len(next) > 1 {
				return nil, fmt.Errorf("build: %s does not support multiple next steps", step.Name)
			} else if len(next) != 0 {
				n = next[0]
			}

			var err error
			next, err = s.Compile(this, n, i)
			if err != nil {
				return nil, fmt.Errorf("cdk step: %s: %w", step.Name, err)
			}
		case types.CdkStepFanOut:
			var err error
			next, err = s.Compile(this, next, i)
			if err != nil {
				return nil, fmt.Errorf("cdk step fanout: %s: %w", step.Name, err)
			}
		default:
			return nil, fmt.Errorf("not a cdk Step: %T: %+v", step.Value, step.Value)
		}

	}

	return next, nil
}

func (a *App) Build(app awscdk.App) error {
	synth := app.Synth(nil)
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

	if os.Getenv("MSH_SKIPDEPLOY") == "" {
		err := Deploy(synth)
		if err != nil {
			return fmt.Errorf("build: failed to deploy: %w", err)
		}
	}

	return nil
}

func Deploy(synth cxapi.CloudAssembly) error {
	cmd := exec.Command("cdk", "bootstrap", "--app", *synth.Directory())
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
