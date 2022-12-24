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
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	lambdaTypes "github.com/aws/aws-sdk-go-v2/service/lambda/types"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/command"
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
		Name:   aws.String("test-4919"),
		Cmd:    cmd,
		client: lambda.NewFromConfig(cfg),
	}

	l.stdoutR, l.Stdout = io.Pipe()

	return l
}

type LambdaCmd struct {
	command.Command
	Name    *string
	lambda  *lambda.GetFunctionOutput
	Stdin   io.ReadCloser
	Stdout  io.WriteCloser
	client  *lambda.Client
	Cmd     command.Command
	stdoutR *io.PipeReader
}

func (s *LambdaCmd) Arn() *string {
	return s.describe().Configuration.FunctionArn
}

func (s *LambdaCmd) describe() *lambda.GetFunctionOutput {
	if s.lambda == nil {
		s.lambda = lo.Must(s.client.GetFunction(context.TODO(), &lambda.GetFunctionInput{
			FunctionName: s.Name,
		}))
	}
	return s.lambda
}

func (s *LambdaCmd) Deploy() error {
	code := lo.Must(LambdaZip())

	_, err := s.client.CreateFunction(context.TODO(), &lambda.CreateFunctionInput{
		Code: &lambdaTypes.FunctionCode{
			ZipFile: code,
		},
		FunctionName: s.Name,
		Handler:      aws.String("lambda_handler.lambda_handler"),
		MemorySize:   aws.Int32(128),
		Publish:      true,
		Role:         aws.String("arn:aws:iam::336983520827:role/service-role/msh-test-role-ztmr8mjk"),
		Runtime:      lambdaTypes.RuntimePython39,
		Timeout:      aws.Int32(900),
	})

	if _, ok := lo.ErrorsAs[*lambdaTypes.ResourceConflictException](err); ok {
		L.Debug.Println("function already exists, updating")
		//lo.Must(s.client.UpdateFunctionCode(context.TODO(), &lambda.UpdateFunctionCodeInput{
		//	Publish:      true,
		//	FunctionName: name,
		//	ZipFile:      code,
		//}))
	} else if err != nil {
		return fmt.Errorf("failed to create function: %w", err)
	}
	L.Debug.Println("function deployed:", *s.Arn())

	return nil
}

func (s *LambdaCmd) Run() error {
	defer s.Stdout.Close()

	stdin := bufio.NewScanner(s.Stdin)

	for stdin.Scan() {

		payload := s.Input(stdin.Text())
		L.Debug.Println("lambda: payload:", string(payload))

		resp := lo.Must(s.client.Invoke(context.TODO(), &lambda.InvokeInput{
			FunctionName:   s.describe().Configuration.FunctionName,
			InvocationType: lambdaTypes.InvocationTypeRequestResponse,
			Payload:        []byte(payload),
			LogType:        lambdaTypes.LogTypeTail,
		}))
		L.Debug.Println("lambda response:", string(resp.Payload))

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
		return nil, fmt.Errorf("failed to write lambda handler: %w", err)
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
	var v []map[string]string
	err := json.Unmarshal(resp, &v)
	if err != nil {
		return nil
	}

	for _, o := range v {
		lo.Must(fmt.Fprintf(pw, o["body"]))
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
