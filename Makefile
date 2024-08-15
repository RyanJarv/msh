export MSH_SKIPDEPLOY=1

.PHONY: build clean

pull:
	git submodule update --init

build: clean
	go generate ./...
	mkdir -p out
	go build -o out/msh cmd/*.go
	./scripts/link_cmds.sh

clean:
	rm -rf out

build-event:
	go build -o out/@cron cmd/cron.go
	go build -o out/@event cmd/event.go

build-filter:
	go build -o out/.filter cmd/sfn.filter.go

build-aws:
	go build -o out/msh.aws cmd/aws.go

build-api:
	go build -o out/msh.api cmd/api.go

build-sleep:
	go build -o out/msh.sleep cmd/sleep.go

install: build
	mkdir -p ~/.msh/bin
	echo 'export PATH=$$HOME/.msh/bin/:$$PATH' > ~/.msh/env
	cp out/msh ~/.msh/bin/msh
	cp test/* ~/.msh/bin/
	chmod +x ~/.msh/bin/msh

test: test-each test-sleep test-lambda
	@echo "\nRunning Tests\n"

test-each:
	@./scripts/test.sh './out/.foreach ./test/mail.json'

test-sleep:
	@./scripts/test.sh 'cat ./test/sfn.json | ./out/.sleep 3'

test-lambda:
	@./scripts/test.sh 'sfn{ ./out/.sleep 34 }'
