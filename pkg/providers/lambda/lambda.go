package lambda

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	lambdaTypes "github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/ryanjarv/msh/pkg/fd"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/command"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io"
	"strings"
)

//go:embed function/lambda_function.py
var lambdaHandler string

func NewLambdaCmd(cmd command.Command) *LambdaCmd {
	cfg := lo.Must(config.LoadDefaultConfig(context.TODO()))

	l := &LambdaCmd{
		Name:   aws.String("msh-function"),
		Cmd:    cmd,
		lambda: lambda.NewFromConfig(cfg),
		iam:    iam.NewFromConfig(cfg),
	}

	l.stdoutR, l.Stdout = io.Pipe()

	return l
}

type LambdaCmd struct {
	command.Command
	Name     *string
	function *lambda.GetFunctionOutput
	Stdin    io.ReadCloser
	Stdout   io.WriteCloser
	lambda   *lambda.Client
	iam      *iam.Client
	Cmd      command.Command
	stdoutR  *io.PipeReader
}

func (s *LambdaCmd) Arn() *string {
	return s.describe().Configuration.FunctionArn
}

func (s *LambdaCmd) describe() *lambda.GetFunctionOutput {
	if s.function == nil {
		s.function = lo.Must(s.lambda.GetFunction(context.TODO(), &lambda.GetFunctionInput{
			FunctionName: s.Name,
		}))
	}
	return s.function
}

func (s *LambdaCmd) Deploy() error {
	role := lo.Must(s.role())

	code := lo.Must(LambdaZip())

	resp, err := s.lambda.GetFunction(context.TODO(), &lambda.GetFunctionInput{FunctionName: s.Name})

	if _, ok := lo.ErrorsAs[*lambdaTypes.ResourceNotFoundException](err); ok {
		_, err = s.lambda.CreateFunction(context.TODO(), &lambda.CreateFunctionInput{
			Code: &lambdaTypes.FunctionCode{
				ZipFile: code,
			},
			FunctionName: s.Name,
			Handler:      aws.String("lambda_handler.lambda_handler"),
			MemorySize:   aws.Int32(128),
			Publish:      true,
			Role:         role.Arn,
			Runtime:      lambdaTypes.RuntimePython39,
			Timeout:      aws.Int32(900),
		})
		if _, ok := lo.ErrorsAs[*lambdaTypes.ResourceConflictException](err); ok {
			L.Debug.Println("function not found, creating")
			lo.Must(s.lambda.UpdateFunctionCode(context.TODO(), &lambda.UpdateFunctionCodeInput{
				Publish:      true,
				FunctionName: s.Name,
				ZipFile:      code,
			}))
		} else if err != nil {
			return fmt.Errorf("failed to create function: %w", err)
		}
		L.Debug.Println("function deployed:", *s.Arn())

	} else if err != nil {
		return fmt.Errorf("failed to get function: %w", err)
	}

	if sha := utils.Sha256(code); *s.describe().Configuration.CodeSha256 != string(sha) {
		L.Debug.Printf("function code changed, updating: (old: %s, new: %s)", *resp.Configuration.CodeSha256, sha)
		lo.Must(s.lambda.UpdateFunctionCode(context.TODO(), &lambda.UpdateFunctionCodeInput{
			FunctionName: s.Name,
			Publish:      true,
			ZipFile:      code,
		}))
	}

	return nil
}

func (s *LambdaCmd) role() (*iamTypes.Role, error) {
	resp, err := s.iam.GetRole(context.TODO(), &iam.GetRoleInput{RoleName: s.Name})
	if _, ok := lo.ErrorsAs[*iamTypes.NoSuchEntityException](err); ok {
		L.Debug.Println("role not found, creating")
		return s.createRole()
	}

	L.Debug.Println("role found:", *resp.Role.Arn)

	return resp.Role, utils.Wrap(err, "failed to get role")
}

func (s *LambdaCmd) createRole() (*iamTypes.Role, error) {
	role := lo.Must(s.iam.CreateRole(context.TODO(), &iam.CreateRoleInput{
		RoleName:                 s.Name,
		AssumeRolePolicyDocument: s.TrustPolicy().ToString(),
	}))
	L.Debug.Println("role created:", *role.Role.Arn)

	lo.Must(s.iam.PutRolePolicy(context.TODO(), &iam.PutRolePolicyInput{
		PolicyDocument: s.LogsPolicy().ToString(),
		PolicyName:     aws.String("lambda"),
		RoleName:       role.Role.RoleName,
	}))
	L.Debug.Println("inline role policy attached to", *role.Role.RoleName)
	return role.Role, nil
}

func (s *LambdaCmd) LogsPolicy() *types.IamPolicy {
	return &types.IamPolicy{
		Name:    fmt.Sprintf("%s-logs", *s.Name),
		Version: "2012-10-17",
		Statement: []types.IamPolicyStatement{
			{
				Effect: "Allow",
				Action: []string{
					"logs:CreateLogGroup",
					"logs:CreateLogStream",
					"logs:PutLogEvents",
				},
				Resource: []string{
					fmt.Sprintf("arn:aws:logs:*:*:log-group:/aws/lambda/%s", *s.Name),
					fmt.Sprintf("arn:aws:logs:*:*:log-group:/aws/lambda/%s:log-stream:*", *s.Name),
				},
			},
		},
	}
}

func (s *LambdaCmd) RunPolicy() *types.IamPolicy {
	return &types.IamPolicy{
		Version: "2012-10-17",
		Name:    fmt.Sprintf("%s-run", *s.Name),
		Statement: []types.IamPolicyStatement{
			{
				Effect: "Allow",
				Action: []string{"lambda:InvokeFunction"},
				Resource: []string{
					*s.Arn(),
				},
			},
		},
	}
}

func (s *LambdaCmd) TrustPolicy() *types.IamPolicy {
	return &types.IamPolicy{
		Name:    fmt.Sprintf("%s-trust", *s.Name),
		Version: "2012-10-17",
		Statement: []types.IamPolicyStatement{
			{
				Effect: "Allow",
				Action: []string{"sts:AssumeRole"},
				Principal: &types.Principal{
					Service: "lambda.amazonaws.com",
				},
			},
		},
	}
}

func (s *LambdaCmd) Run() error {
	defer s.Stdout.Close()

	stdin := bufio.NewScanner(s.Stdin)

	for stdin.Scan() {

		body := lo.Must(json.Marshal(fd.Event{
			Type:    fd.MessageEvent,
			Content: stdin.Text(),
			Id:      0,
		}))
		payload := fmt.Sprintf("[%s]", s.Input(string(body)))
		L.Debug.Println("function: payload:", payload)

		resp := lo.Must(s.lambda.Invoke(context.TODO(), &lambda.InvokeInput{
			FunctionName:   s.describe().Configuration.FunctionName,
			InvocationType: lambdaTypes.InvocationTypeRequestResponse,
			Payload:        []byte(payload),
			LogType:        lambdaTypes.LogTypeTail,
		}))
		L.Debug.Println("function response:", string(resp.Payload))

		err := WriteOutput(s.Stdout, resp.Payload)
		if err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}

		L.Debug.Println("function logs:", Base64Decode(*resp.LogResult)+"\n")
	}

	return nil
}

func (s *LambdaCmd) SetStdin(p interface{}) {
	s.Stdin = p.(io.ReadCloser)
}

func (s *LambdaCmd) GetStdout() io.Reader {
	return s.stdoutR
}

func LambdaZip() ([]byte, error) {
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	f, err := w.Create("lambda_handler.py")
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	_, err = f.Write([]byte(lambdaHandler))
	if err != nil {
		return nil, fmt.Errorf("failed to write function handler: %w", err)
	}

	err = w.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close zip writer: %w", err)
	}

	return io.ReadAll(buf)
}

func Base64Decode(s string) string {
	return string(lo.Must(base64.StdEncoding.DecodeString(s)))
}

func WriteOutput(pw io.Writer, resp []byte) []string {
	var msgs []string
	err := json.Unmarshal(resp, &msgs)
	if err != nil {
		L.Error.Fatalln("unmarshal response to string:", err)
	}

	for _, o := range msgs {
		event := fd.Event{}
		err = json.Unmarshal([]byte(o), &event)
		if err != nil {
			L.Error.Fatalln("unmarshal response to map:", err)
		}

		lo.Must(fmt.Fprintf(pw, event.Content))
	}

	return nil
}

type InputTemplate struct {
	Event string            `json:"event"`
	Cmd   []string          `json:"cmd"`
	Env   map[string]string `json:"env"`
}

func (s *LambdaCmd) Input(body string) string {
	env := map[string]string{}

	if s.Cmd.Env != nil {
		for _, envVar := range s.Env {
			parts := strings.SplitN(envVar, "=", 2)
			env[parts[0]] = parts[1]
		}
	}

	input := lo.Must(utils.JSONMarshal(map[string]interface{}{
		"body": body,
		"cmd":  s.Cmd.Args,
		"env":  env,
	}))

	return string(input)
}
