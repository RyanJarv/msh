package msh

import (
	"bufio"
	"context"
	"github.com/docker/go-connections/nat"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func NewLambdaExec(g Global) LambdaExec {
	return LambdaExec{
		Aws:       g.Aws,
		ctx:       context.Background(),
		Container: NewContainer(NewDockerInput{
			Global: g,
			Tag: "latest",
		}),
	}
}

type LambdaExec struct {
	Container
	Aws 	  Aws
	ctx       context.Context
	Name      string
	remoteUri string
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

func (l *LambdaExec) Start() {
	l.Container.Start(
		nat.PortSet{ LambdaRiePort: {} },
		map[nat.Port][]nat.PortBinding{ LambdaRiePort: {
			{
				HostIP:   "127.0.0.1",
				HostPort: "0", // Picks a random port
			},
		}},
	)
}
