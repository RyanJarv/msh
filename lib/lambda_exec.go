package lib

import (
	"bufio"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func NewLambdaExec(imageName string) LambdaExec {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return LambdaExec{
		ctx:      context.Background(),
		cli:      cli,
		localTag: imageName,
	}
}

type LambdaExec struct {
	Cmd           *exec.Cmd
	cli           *client.Client
	ctx           context.Context
	ID            string
	PortBinding   nat.PortBinding
	ContainerName string
	localTag      string
	repo          string
	remoteUri     string
}

func (l LambdaExec) StdinToHttp(url url.URL) {
	go func() {
		input := bufio.NewScanner(os.Stdin)
		client := http.DefaultClient
		for input.Scan() {
			t := input.Text()
			post, err := client.Post(url.String(), "application/json", strings.NewReader(t))
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stderr, post.Body)
		}
		if err := input.Err(); err != nil {
			panic(err)
		}
	}()
	return
}

func (l LambdaExec) BuildContainer(buildContext io.Reader) {
	build, err := l.cli.ImageBuild(l.ctx, buildContext, types.ImageBuildOptions{
		Tags: []string{l.localTag},
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

func (l LambdaExec) ImageLoad(err error, build types.ImageBuildResponse) []byte {
	load, err := l.cli.ImageLoad(l.ctx, build.Body, false)
	defer load.Body.Close()

	all, err := ioutil.ReadAll(load.Body)
	if err != nil {
		panic(err)
	}
	return all
}

// ContainerLogs attaches to the container, if logs is set it will display logs since the beginning of execution.
func (l *LambdaExec) ContainerLogs() {
	attach, err := l.cli.ContainerLogs(l.ctx, l.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stderr, attach)
}

func (l LambdaExec) ContainerStop() {
	l.cli.ContainerStop(l.ctx, l.ID, nil)
	if err := l.cli.ContainerRemove(l.ctx, l.ID, types.ContainerRemoveOptions{}); err != nil {
		fmt.Printf("[ERROR] unable to remove container %s\n", l.ID)
	}
}

func (l LambdaExec) ContainerWait() {
	statusCh, errCh := l.cli.ContainerWait(l.ctx, l.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}
}

func (l *LambdaExec) ContainerStart() {
	if resp, err := l.containerCreate(); err != nil {
		panic(err)
	} else {
		l.ID = resp.ID

		for _, warn := range resp.Warnings {
			fmt.Fprint(os.Stderr, warn)
		}
	}

	if err := l.cli.ContainerStart(l.ctx, l.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	if inspect, err := l.cli.ContainerInspect(l.ctx, l.ID); err != nil {
		panic(err)
	} else {
		l.PortBinding = inspect.NetworkSettings.NetworkSettingsBase.Ports[LambdaRiePort][0]
		l.ContainerName = inspect.Name
	}
}

func (l *LambdaExec) containerCreate() (container.ContainerCreateCreatedBody, error) {
	resp, err := l.cli.ContainerCreate(l.ctx, &container.Config{
		Image: l.localTag,
		Tty:   false,
		ExposedPorts: nat.PortSet{
			"8080/tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{
			"8080/tcp": {
				{
					HostIP:   "127.0.0.1",
					HostPort: "0", // Picks a random port
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		panic(errors.Wrap(err, "starting container"))
	}

	return resp, err
}
