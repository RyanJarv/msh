export MSH_SKIPDEPLOY=1


build: build-event build-sleep build-api build-aws build-filter
	go generate ./...
	go build -o out/sfn cmd/sfn.go
	go build -o out/@sns cmd/sns.go
	go build -o out/.map cmd/sfn.map.go
	go build -o out/.each cmd/each.go
	go build -o out/.mail cmd/mail.go
	go build -o out/.lambda.python cmd/lambda.python.go

build-event:
	go build -o out/@cron cmd/cron.go
	go build -o out/@event cmd/event.go

build-filter:
	go build -o out/.filter cmd/sfn.filter.go

build-aws:
	go build -o out/.aws cmd/aws.go

build-api:
	go build -o out/.api cmd/api.go

build-sleep:
	go build -o out/.sleep cmd/sleep.go

install: build
	mkdir -p ~/.msh/bin
	echo 'export PATH=$$HOME/.msh/bin/:$$PATH' > ~/.msh/env
	cp out/msh ~/.msh/bin/msh
	cp test/* ~/.msh/bin/
	chmod +x ~/.msh/bin/msh

test: build test-each test-sleep test-lambda
	@echo "\nRunning Tests\n"

test-each:
	@./scripts/test.sh './out/each ./test/mail.json'

test-sleep:
	@./scripts/test.sh 'cat ./test/sfn.json | ./out/sleep 3'

test-lambda:
	@./scripts/test.sh './out/lambda.python ./test/test.py | ./out/sleep 34'
