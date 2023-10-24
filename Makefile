build:
	go generate ./...
	go build -o out/@event cmd/event.go
	go build -o out/sfn cmd/sfn.go
	go build -o out/sns cmd/sns.go
	go build -o out/mail cmd/mail.go
	go build -o out/lambda.python cmd/lambda.python.go
	go build -o out/sleep cmd/sleep.go

docker-build:
	docker build -f test/python.msh -t python.msh .

install: build
	mkdir -p ~/.msh/bin
	echo 'export PATH=$$HOME/.msh/bin/:$$PATH' > ~/.msh/env
	cp out/msh ~/.msh/bin/msh
	cp test/* ~/.msh/bin/
	chmod +x ~/.msh/bin/msh

# Using different buffers for the config and data can cause the data not to be read.
test-stress:
	for i in $$(seq 1 10); do ./scripts/test.sh './out/msh.local cat ./test/data/test | ./out/msh.local cat' 'test'; done

test-stress-sqs:
	for i in $$(seq 1 10); do ./scripts/test.sh 'echo "input"| ./out/sqs' 'input'; done

test-sessions:
	for i in $$(seq 1 10 ); do echo "input"| ./out/msh.pipes cat | cat; done
	for i in $$(seq 1 10); do ./scripts/test.sh 'echo "input"| ./out/msh.pipes cat' 'input'; done

test-local:
	@echo "\nRunning Local Tests\n"

	@./scripts/test.sh './out/msh.local cat ./test/data/test | ./out/msh.local cat' 'test'

	@# Not sure why this doesn't work when running as a test
	@#./scripts/test.sh './out/msh.local cat ./test/data/sqs-stdin | ./out/sqs | ./out/msh.local cat' 'input'

test-pass-by-ref:
	@./scripts/test.sh 'echo "input"| ./out/msh.pipes sed s/input/output/ | cat' 'output' 'Url' 'input'

test: build test-local test-stress
	@echo "\nRunning Tests\n"

	@go run ./cmd/msh.lambda.go ./test/test.py | go run ./cmd/msh.sleep.go 34 

	# Old stuff:
	@./scripts/test.sh './out/msh.local echo "input"| ./out/sqs ' 'input'

	@./scripts/test.sh 'echo "input"| ./out/sqs' 'input'

	@./scripts/test.sh 'echo "input"| ./out/msh.lambda sed s/input/output/' 'output'
	@./scripts/test.sh 'echo "input"| ./out/msh.pipes sed s/input/output/' 'output'
	@./scripts/test.sh 'cat test/data/input | ./out/msh.lambda sed "s/input/output/"' 'output'
	@./scripts/test.sh 'cat test/data/input | ./out/msh.pipes sed "s/input/output/"' 'output'
	@./scripts/test.sh 'cat test/data/input | ./out/msh.lambda sed "s/input/output/"' 'output'
	@#./scripts/test.sh 'echo "input"| ./out/sqs | ./out/msh.local bash ./test/cat.msh.sqs | ./out/sqs' 'output'
	@#./scripts/test.sh 'echo "input"| ./out/sqs | ./out/msh.local bash ./test/cat.msh.sqs | ./out/sqs' 'output'
	@#./scripts/test.sh 'echo "input"|./out/msh.lambda' 'output'
