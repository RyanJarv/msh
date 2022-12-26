package eventbridge

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/pipes"
	pipesTypes "github.com/aws/aws-sdk-go-v2/service/pipes/types"
	"github.com/ryanjarv/msh/pkg/fd"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/lambda"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io"
	"strings"
)

func NewPipe(lambda *lambda.LambdaCmd) AwsPipe {
	cfg := lo.Must(config.LoadDefaultConfig(context.TODO()))

	pipe := AwsPipe{
		CreatePipeInput: &pipes.CreatePipeInput{
			EnrichmentParameters: &pipesTypes.PipeEnrichmentParameters{
				InputTemplate: aws.String(lambda.Input("<$.body>")),
			},
			DesiredState: pipesTypes.RequestedPipeStateRunning,
		},
		LambdaCmd: lambda,
		cfg:       lo.Must(config.LoadDefaultConfig(context.TODO())),
		events:    pipes.NewFromConfig(cfg),
		iam:       iam.NewFromConfig(cfg),
	}

	return pipe
}

type AwsPipe struct {
	CreatePipeInput *pipes.CreatePipeInput
	cfg             aws.Config
	LambdaCmd       *lambda.LambdaCmd
	events          *pipes.Client
	iam             *iam.Client
	iamPolicies     []*types.IamPolicy
	Stdout          fd.ISqs
	Stdin           fd.ISqs
}

func (p *AwsPipe) Name() string {
	hash := utils.Sha256Base64(
		strings.Join(p.LambdaCmd.Cmd.Args, ""),
	)
	hash = strings.ReplaceAll(hash, "=", "")

	l := lo.Min([]int{len(hash), 8})
	name := fmt.Sprintf("msh-%s-%s", p.LambdaCmd.Cmd.Args[0], hash[0:l])

	return name
}

func (p *AwsPipe) Deploy() error {
	err := p.LambdaCmd.Deploy()
	if err != nil {
		return fmt.Errorf("AwsPipe deploy: %w", err)
	}

	p.CreatePipeInput.Name = aws.String(p.Name())

	p.CreatePipeInput.Source = p.Stdin.Arn()
	p.CreatePipeInput.SourceParameters = p.Stdin.PipeSourceParameters()
	p.CreatePipeInput.Enrichment = p.LambdaCmd.Arn()
	p.CreatePipeInput.Target = p.Stdout.Arn()

	err = p.DeployIAM()
	if err != nil {
		return fmt.Errorf("deploy: %w", err)
	}

	_, err = p.events.CreatePipe(context.TODO(), p.CreatePipeInput)
	if _, ok := lo.ErrorsAs[*pipesTypes.ConflictException](err); ok {
		L.Debug.Println("reusing existing pipe:", p.Name())
	} else if err != nil {
		return fmt.Errorf("failed to create pipe: %w\n", err)
	}

	return nil
}

func (p *AwsPipe) Run() error {
	p.Stdout.Wait()
	return nil
}

func (p *AwsPipe) SetStdin(source interface{}) {
	p.Stdin = lo.Must(fd.NewSqsFrom(context.TODO(), source, p.Name(), "stdin"))
	L.Debug.Println("eventbridge pipes: source:", p.Stdin.Arn())
	//p.LambdaCmd.SetStdin(p.Stdin)
}

func (p *AwsPipe) GetStdout() io.ReadCloser {
	cfg := lo.Must(config.LoadDefaultConfig(context.TODO()))
	p.Stdout = lo.Must(fd.CreateSqs(cfg, p.Name(), "stdout"))
	L.Debug.Println("eventbridge pipes: target:", p.Stdout.Arn())

	return p.Stdout
}

func (p *AwsPipe) DeployIAM() error {
	p.addPolicy(p.LambdaCmd.RunPolicy())
	p.addPolicy(p.Stdin.ReadPolicy())
	p.addPolicy(p.Stdout.WritePolicy())

	var role *iamTypes.Role

	resp, err := p.iam.GetRole(context.TODO(), &iam.GetRoleInput{RoleName: aws.String(p.Name())})
	if _, ok := lo.ErrorsAs[*iamTypes.NoSuchEntityException](err); ok {
		L.Debug.Println("role not found, creating")

		role, err = p.createRole()
		if err != nil {
			return fmt.Errorf("failed to create role: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to get role: %w", err)
	} else {
		role = resp.Role
	}

	L.Debug.Println("role found:", *role.Arn)

	for _, policy := range p.iamPolicies {
		_, err := p.iam.PutRolePolicy(context.TODO(), &iam.PutRolePolicyInput{
			RoleName:       role.RoleName,
			PolicyName:     aws.String(policy.Name),
			PolicyDocument: policy.ToString(),
		})
		if err != nil {
			return fmt.Errorf("failed to put role policy: %w", err)
		}

		L.Debug.Printf("attached %s inline policy to %s\n", policy.Name, *role.Arn)
	}

	p.CreatePipeInput.RoleArn = role.Arn

	return utils.Wrap(err, "failed to get role")
}

func (p *AwsPipe) createRole() (*iamTypes.Role, error) {
	role := lo.Must(p.iam.CreateRole(context.TODO(), &iam.CreateRoleInput{
		RoleName:                 aws.String(p.Name()),
		AssumeRolePolicyDocument: p.TrustPolicy().ToString(),
	}))
	L.Debug.Println("role created:", *role.Role.Arn)

	return role.Role, nil
}

func (p *AwsPipe) TrustPolicy() *types.IamPolicy {
	return &types.IamPolicy{
		Name:    fmt.Sprintf("%s-trust", p.Name()),
		Version: "2012-10-17",
		Statement: []types.IamPolicyStatement{
			{
				Effect: "Allow",
				Action: []string{"sts:AssumeRole"},
				Principal: &types.Principal{
					Service: "pipes.amazonaws.com",
				},
			},
		},
	}
}

func (p *AwsPipe) addPolicy(policy *types.IamPolicy) {
	p.iamPolicies = append(p.iamPolicies, policy)
}
