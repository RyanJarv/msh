build:
	go generate ./...
	go build -o out/msh main.go

install: build
	cp out/msh /usr/local/bin/msh
	chmod +x /usr/local/bin/msh

test: build dockerfile/echo dockerfile/cat compose/echo

test/remote: ecs/echo

dockerfile/echo: build
	./scripts/test.sh './test/dockerfile/echo hello from args' 'hello .* args'

dockerfile/cat: build
	echo 'echo "hello from a pipe"' | ./scripts/test.sh ./test/dockerfile/cat 'hello .* pipe'

ecs/echo: build
	./scripts/test.sh 'msh ecs ./test/dockerfile/echo hello from ecs' 'hello .* ecs'

compose/echo: build
	./scripts/test.sh './test/compose/echo hello from echo compose' 'hello .* echo compose'

compose/cat: build
	echo 'echo "hello from a pipe to cat compose"' | ./scripts/test.sh ./test/compose/cat 'hello .* pipe .* cat compose'

compose/lambda: build
	./scripts/env.sh ./test/compose/lambda/compose.yaml app.lambda_handler & sleep 2
	./scripts/test.sh 'curl -s -XPOST http://localhost:9000/2015-03-31/functions/function/invocations -d {}' 'hello .* lambda compose'
