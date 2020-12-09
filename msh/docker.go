package msh

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	ecrTypes "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	L "github.com/ryanjarv/msh/logger"
	"io"
	"os"
	"strings"
)

type NewDockerInput struct {
	Tag    string
	Global Global
}

func NewContainer(input NewDockerInput) Container {
	c, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	return Container{
		NewDockerInput: input,
		client: c,
		ecr: ecr.NewFromConfig(input.Global.Aws.Config),
		ctx: context.Background(),
	}
}

type Container struct {
	NewDockerInput
	client      *client.Client
	ecr   		*ecr.Client
	ctx         context.Context
	ImageID     string
	Inspect     types.ContainerJSON
	PortBinding map[nat.Port][]nat.PortBinding
	Name        string
}

func (d *Container) Build(buildContext io.Reader) {
	build, err := d.client.ImageBuild(d.ctx, buildContext, types.ImageBuildOptions{
		Tags: []string{d.Tag},
		Dockerfile: "Dockerfile",
		PullParent: true,
		SuppressOutput: true,
	})
	if err != nil {
		panic(err)
	}
	defer build.Body.Close()
	io.Copy(os.Stderr, build.Body)
}

func (d *Container) Stop() {
	d.client.ContainerStop(d.ctx, d.ImageID, nil)
	if err := d.client.ContainerRemove(d.ctx, d.ImageID, types.ContainerRemoveOptions{}); err != nil {
		fmt.Printf("[ERROR] unable to remove container %s\n", d.ImageID)
	}
}


func (d Container) Wait() {
   statusCh, errCh := d.client.ContainerWait(d.ctx, d.ImageID, container.WaitConditionNotRunning)
   select {
   case err := <-errCh:
	   if err != nil {
			panic(err)
	   }
   case <-statusCh:
   }
}


func (d *Container) Start(portSet nat.PortSet, portBindings map[nat.Port][]nat.PortBinding) {
	d.PortBinding = portBindings
	d.create(portSet, portBindings)

	if err := d.client.ContainerStart(d.ctx, d.ImageID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	if inspect, err := d.client.ContainerInspect(d.ctx, d.ImageID); err != nil {
		panic(err)
	} else {
		d.Inspect = inspect
		d.Name = inspect.Name
	}
}

func (d *Container) create(portSet nat.PortSet, portBindings map[nat.Port][]nat.PortBinding) container.ContainerCreateCreatedBody {
	resp, err := d.client.ContainerCreate(d.ctx,
		&container.Config{
			Image:        d.Tag,
			Tty:          false,
			ExposedPorts: portSet,
		},
		&container.HostConfig{ PortBindings: portBindings },
	 nil, nil, "")
	if err != nil {
		panic(errors.Wrap(err, "starting container"))
	}
	d.ImageID = resp.ID
	for _, warn := range resp.Warnings {
		L.Debug.Println(os.Stderr, warn)
	}
	return resp
}

// ContainerLogs attaches to the container, if logs is set it will display logs since the beginning of execution.
func (d *Container) Logs() {
	attach, err := d.client.ContainerLogs(d.ctx, d.ImageID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stderr, attach)
}

func (d *Container) NewTag(tag string) {
	d.client.ImageTag(d.ctx, d.Tag, tag)
}

func (d *Container) Push(ecrClient ecr.Client, remoteUri string) {
	token, err := ecrClient.GetAuthorizationToken(d.ctx, &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		panic(err)
	}

	// SDK seems to encode it in the wrong format so, decode and reencode
	auth := types.AuthConfig{
		Username: "AWS",
	} // UserName is always AWS

	if s, err := base64.StdEncoding.DecodeString(*token.AuthorizationData[0].AuthorizationToken); err != nil {
		panic(err)
	} else {
		auth.Password = strings.Split(string(s), ":")[1] // user:pass
	}
	encodedJSON, err := json.Marshal(auth)
	if err != nil {
		panic(err)
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	push, err := d.client.ImagePush(d.ctx, remoteUri, types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		panic(err)
	}
	defer push.Close()

	_, err = io.Copy(L.Debug.Writer(), push)
}

func (d *Container) CreateRepo(repoPath string) {
	_, err := d.ecr.CreateRepository(d.ctx, &ecr.CreateRepositoryInput{RepositoryName: aws.String(repoPath)})
	var exists *ecrTypes.RepositoryAlreadyExistsException
	if err != nil && !errors.As(err, &exists) {
		panic(err)
	}
}

