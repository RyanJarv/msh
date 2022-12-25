build:
	go generate ./...
	go build -o out/msh.local cmd/msh.local.go
	go build -o out/msh.lambda cmd/msh.lambda.go
	go build -o out/msh.pipes cmd/msh.pipes.go
	go build -o out/sqs cmd/sqs.go

docker-build:
	docker build -f test/python.msh -t python.msh .

install: build
	mkdir -p ~/.msh/bin
	echo 'export PATH=$$HOME/.msh/bin/:$$PATH' > ~/.msh/env
	cp out/msh ~/.msh/bin/msh
	cp test/* ~/.msh/bin/
	chmod +x ~/.msh/bin/msh

# Using different buffers for the config and data can cause the data not to be read.
test-buffers:
	for i in $$(seq 1 10); do ./scripts/test.sh './out/msh.local cat ./test/data/test | ./out/msh.local cat' 'test'; done

test-local:
	@echo "\nRunning Local Tests\n"

	@./scripts/test.sh './out/msh.local cat ./test/data/test | ./out/msh.local cat' 'test'

	@# Not sure why this doesn't work when running as a test
	@#./scripts/test.sh './out/msh.local cat ./test/data/sqs-stdin | ./out/sqs | ./out/msh.local cat' 'input'

test: build test-local test-buffers
	@echo "\nRunning Tests\n"

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
