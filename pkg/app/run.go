package app

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
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

func (a *App) Run(step types.CdkStep) error {
	a.State.AddStep(step)

	if !utils.IsTTY(a.Stdout) {
		return a.State.WriteState(a.Stdout)
	}

	if utils.IsTTY(a.Stdout) || os.Getenv("MSH_BUILD") == "always" {
		app := awscdk.NewApp(&awscdk.AppProps{})
		stack := awscdk.NewStack(app, jsii.String("msh"), &awscdk.StackProps{})

		if _, err := a.Compile(stack, nil, 0); err != nil {
			return err
		}

		return a.Build(app)
	}

	return nil
}

func (s *State) WriteState(f *os.File) error {
	d, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("writeState: failed to marshal app: %w", err)
	}

	L.Debug.Printf("State.WriteState: writing to %s: %s", f.Name(), string(d))
	_, err = io.WriteString(f, string(d)+"\n")
	if err != nil {
		return fmt.Errorf("writeState: failed to write app: %w", err)
	}

	return nil
}

func (a *App) Compile(scope constructs.Construct, next types.CdkStep, i int) (types.CdkStep, error) {
	L.Debug.Println("App.Compile")
	// Reverse and build backwards so that Iterator, Next, AddTarget calls receive the proceeding target CdkStep.
	steps := lo.Reverse(a.State.Steps)

	//var next types.CdkStep

	for i, step := range steps {
		scope := constructs.NewConstruct(scope, jsii.String(fmt.Sprintf("%s%d", step.Name, i)))

		cdkstep, ok := step.Value.(types.CdkStep)
		if !ok {
			return nil, fmt.Errorf("step must implement CdkStep, got: %T", step.Value)
		}

		// Pass the index of the command from the start of the pipeline starting from 1. This
		// is used to ensure unique naming of resources in the CDK.
		cmdNum := len(steps) - i + 1
		err := cdkstep.Compile(scope, cmdNum)
		if err != nil {
			return nil, fmt.Errorf("compile: %w", err)
		}

		switch s := cdkstep.(type) {
		// We don't use Next on iterators currently, so make sure this is higher priority than INextable.
		case types.IIterator:
			chain, err := EnsureNext(scope, next)
			if err != nil {
				return nil, err
			}

			s.Iterator(chain.StartState())
		case sfn.INextable:
			chain, err := EnsureNext(scope, next)
			if err != nil {
				return nil, err
			}

			// Call StartState() to ensure we are not passing the whole CdkStep object which confuses CDK.
			s.Next(chain.StartState())
		case awsevents.Rule:
			n, ok := next.(awsevents.IRuleTarget)
			if !ok {
				return nil, fmt.Errorf("next step must be a event rule target, got: %T", next)
			}

			if next == nil {
				return nil, fmt.Errorf("events rule does not work at the end of a chain")
			}

			s.AddTarget(n)
		case awssns.ITopic:
			n, ok := next.(awssns.ITopicSubscription)
			if !ok {
				return nil, fmt.Errorf("next step must be a topic subscription, got: %T", next)
			}

			if next != nil {
				s.AddSubscription(n)
			}
		default:
			return nil, fmt.Errorf("unknown type for step, got: %T", cdkstep)
		}

		next = cdkstep
	}

	return next, nil
}

func EnsureNext(scope constructs.Construct, next types.CdkStep) (chain sfn.IChainable, err error) {
	if next == nil {
		chain = sfn.NewSucceed(scope, jsii.String("succeed"), &sfn.SucceedProps{})
	} else {
		var ok bool
		chain, ok = next.(sfn.IChainable)
		if !ok {
			return nil, fmt.Errorf("next step must be IChainable, got: %T", next)
		}
	}
	return chain, nil
}

func (a *App) Build(app awscdk.App) error {
	L.Debug.Println("App.Build")
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
