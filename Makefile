build:
	go generate ./...
	go build -o out/msh main.go

install: build
	cp out/msh /usr/local/bin/msh
	chmod +x /usr/local/bin/msh

test: build dockerfile/echo dockerfile/cat ecs/echo

dockerfile/echo: build
	./scripts/test.sh './test/dockerfile/echo hello from args' 'hello .* args'

dockerfile/cat: build
	echo 'echo "hello from a pipe"' | ./scripts/test.sh ./test/dockerfile/cat 'hello .* pipe'

ecs/echo: build
	./scripts/test.sh 'msh ecs ./test/dockerfile/echo hello from ecs' 'hello .* ecs'