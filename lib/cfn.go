package lib

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/pkg/errors"
	L "github.com/ryanjarv/msh/logger"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"
)

type Cfn struct {
	Parser

	// For now Cfn only accepts Content, not Parser.Path
	Content string

	client         *cloudformation.Client
	template       *template.Template
	parsedTemplate *string
	CustomData     map[string]string
	StackOutputs   []types.Output
	StackId        string
	stack          types.Stack
}

//
func NewCfn(cfg Parser) RunTemplate {
	cfg.Aws.Identity = AwsIdentityTmpl(cfg)
	if cfg.Path != "" {
		panic("not implemented")
	}
	return &Cfn{
		Parser: cfg,
		Content: cfg.Content,
		client: cloudformation.NewFromConfig(cfg.Aws.Config),
	}
}

//func (c *Cfn) Parse() Runnable {
//	if str, err := ReadConfig(c.Path); err == nil {
//		c.template, err = template.New("cfn").Parse(str)
//	} else if err != nil {
//		panic(err)
//	}
//	return c.Parser
//}
//
// TODO: After template creation
//if err := CfnLint(path); err != nil {
//	return err
//}

func (c *Cfn) Parse(data map[string]string) {
	var err error
	c.template, err = template.New("cfn").Parse(c.Content)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	c.CustomData = data
	if err := c.template.ExecuteTemplate(&buf, "cfn", c); err != nil {
		panic(err)
	}

	parsed, err := ioutil.ReadAll(&buf)
	if err != nil {
		panic(err)
	}

	c.parsedTemplate = aws.String(StripShebang(string(parsed)))
	return
}

func (c *Cfn) Run() error {
	if c.parsedTemplate == nil {
		c.Parse(map[string]string{})
	}
	L.Debug.Println("\n" + *c.parsedTemplate)

	// Doesn't support docker lambda currently
	//err := CfnLint(strings.NewReader(*c.parsedTemplate))
	//if err != nil {
	//	return err
	//}

	if _, err := c.client.ValidateTemplate(c.Context, &cloudformation.ValidateTemplateInput{
		TemplateBody: c.parsedTemplate,
	}); err != nil {
		panic(err)
	}

	//stack, err := c.Describe()
	//if err != nil {
	//	panic(err)
	//}


	stackInput := &cloudformation.CreateStackInput{
		StackName:        aws.String(c.Project),
		Capabilities:     []types.Capability{types.CapabilityCapabilityIam, types.CapabilityCapabilityAutoExpand},
		OnFailure:        types.OnFailureDoNothing,
		TemplateBody:     c.parsedTemplate,
		TimeoutInMinutes: aws.Int32(5),
	}
	L.Debug.Printf("StackName: %v", *stackInput.StackName)

	out, err := c.client.CreateStack(c.Context, stackInput)
	if err != nil {
		var target *types.AlreadyExistsException
		if errors.As(err, &target) {
			if err := c.Update(); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		c.StackId = *out.StackId
		L.Debug.Printf("Created Stack ID: %s", c.StackId)
	}

Watch:
	for {
		time.Sleep(time.Second * 5)
		c.stack = c.Describe()
		printStackStatus(c.stack)
		switch status := c.stack.StackStatus; status {
		case types.StackStatusCreateComplete, types.StackStatusCreateFailed, types.StackStatusDeleteComplete, types.StackStatusUpdateComplete, types.StackStatusUpdateRollbackFailed:
			c.StackOutputs = c.stack.Outputs
			break Watch
		default:
			printStackStatus(c.stack)
			fmt.Printf(".")
		}
	}

	return nil
}

func (c Cfn) Update() error {
	out, err := c.client.GetTemplate(c.Context, &cloudformation.GetTemplateInput{
		StackName:     aws.String(c.Project),
		TemplateStage: types.TemplateStageOriginal,
	})
	if err != nil {
		panic(err)
	}
	if *out.TemplateBody == *c.parsedTemplate {
		L.Debug.Println("stack up to date")
		return nil
	}
	_, err = c.client.UpdateStack(c.Context, &cloudformation.UpdateStackInput{
		StackName:    aws.String(c.Project),
		Capabilities: []types.Capability{types.CapabilityCapabilityIam, types.CapabilityCapabilityAutoExpand},
		TemplateBody: c.parsedTemplate,
	})
	if err != nil {
		panic(err)
	}
	return err
}

func (c Cfn) Describe() types.Stack {
	stacks, err := c.client.DescribeStacks(c.Context, &cloudformation.DescribeStacksInput{
		StackName: aws.String(c.Project),
	})
	if err != nil {
		panic(err)
	}
	return stacks.Stacks[0]
}

func printStackStatus(stack types.Stack) {
	if stack.StackStatusReason != nil {
		L.Debug.Printf("%s: %s -- %s", *stack.StackName, stack.StackStatus, *stack.StackStatusReason)
	} else {
		L.Debug.Printf("%s: %s", *stack.StackName, stack.StackStatus)
	}
}

func CfnLint(in io.Reader) error {
	var cmd *exec.Cmd
	if abs, err := filepath.Abs("/Users/jarv/Code/msh-repo/conf/cfn-lint/cfn-lint"); err != nil {
		return err
	} else {
		cmd = exec.Command(abs, "-")
	}
	cmd.Stdin = in
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	log.Printf("Running command and waiting for it to finish...")
	return cmd.Run()
}

