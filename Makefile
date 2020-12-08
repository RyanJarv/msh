build:
	go generate ./...
	go build -o out/msh main.go

install: build
	cp out/msh /usr/local/bin/msh
	chmod +x /usr/local/bin/msh

cmd/echo: build
	PATH="${PWD}/out:${PATH}" DEBUG=true ./config/cmds/echo hello world from the args

cmd/cat: build
	echo "hello world from a pipe" | PATH="${PWD}/out:${PATH}" DEBUG=true ./config/cmds/cat

#TODO: Mount lint file
cmd/cfn-lint: build
	PATH="${PWD}/out:${PATH}" DEBUG=true ./config/cmds/cfn-lint ./config/swf/hello_world/iam.yaml --info

ecs: build
	PATH="${PWD}/out:${PATH}" DEBUG=true msh ecs ./config/cmds/echo hello from ecs

swf/iam: build
	PATH="${PWD}/out:${PATH}" DEBUG=true ./config/swf/hello_world/iam.yaml

swf/states: build
	PATH="${PWD}/out:${PATH}" DEBUG=true ./config/swf/hello_world/states.json

