package lambda

import (
	"archive/zip"
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
	"github.com/samber/lo"
	"io"
)

//go:embed function/lambda_function.py
var lambdaHandler string

func NewLambdaCmd(cmd command.Command) *LambdaCmd {
	cfg := lo.Must(config.LoadDefaultConfig(context.TODO()))
	return &LambdaCmd{
		Cmd:    cmd,
		client: lambda.NewFromConfig(cfg),
	}
}

type LambdaCmd struct {
	command.Command
	Lambda *lambda.GetFunctionOutput
	Input  []byte
	pw     *io.PipeWriter
	pr     *io.PipeReader
	client *lambda.Client
	Cmd    command.Command
}

func (s *LambdaCmd) Deploy() error {
	name := aws.String("test-4919")
	code := lo.Must(LambdaZip())

	_, err := s.client.CreateFunction(context.TODO(), &lambda.CreateFunctionInput{
		Code: &lambdaTypes.FunctionCode{
			ZipFile: code,
		},
		FunctionName: name,
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

	s.Lambda = lo.Must(s.client.GetFunction(context.TODO(), &lambda.GetFunctionInput{
		FunctionName: name,
	}))

	return nil
}

func (s *LambdaCmd) Run() error {
	defer s.pw.Close()
	//s.CreatePipeInput.Enrichment = step.Arn()

	resp := lo.Must(s.client.Invoke(context.TODO(), &lambda.InvokeInput{
		FunctionName:   s.Lambda.Configuration.FunctionName,
		InvocationType: lambdaTypes.InvocationTypeRequestResponse,
		Payload:        s.Input,
		LogType:        lambdaTypes.LogTypeTail,
	}))
	L.Debug.Println("lambda response:", string(resp.Payload))

	err := WriteOutput(s.pw, resp.Payload)
	if err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	L.Debug.Println("function logs:", Base64Decode(*resp.LogResult)+"\n")

	return nil
}

func (s *LambdaCmd) SetStdin(p interface{}) {
	stdin := lo.Must(io.ReadAll(p.(io.Reader)))
	s.Input = Input(s.Cmd.Args, stdin)
}

func (s *LambdaCmd) GetStdout() io.Reader {
	s.pr, s.pw = io.Pipe()
	return s.pr
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

func WriteOutput(pw *io.PipeWriter, resp []byte) []string {
	var v []map[string]string
	err := json.Unmarshal(resp, &v)
	if err != nil {
		return nil
	}

	for _, o := range v {
		lo.Must(fmt.Fprintf(pw, o["Content"]))
	}

	return nil
}

func Input(cmd []string, stdin []byte) []byte {
	body := string(lo.Must(json.Marshal(map[string]interface{}{
		"Content": string(stdin),
	})))

	input := []map[string]interface{}{
		{
			"cmd":  cmd,
			"body": body,
		},
	}
	L.Debug.Println("lambda input:", input)
	return lo.Must(json.Marshal(input))
}
