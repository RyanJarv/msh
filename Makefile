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

test: build
	@echo "\nRunning Tests\n"
	#@./scripts/test.sh './test/python.msh --version' 'Python 3.*'
	#@./scripts/test.sh './out/msh.local cat ./test/data/test | cat' 'StdinArn'
	#@./scripts/test.sh './out/msh.local cat ./test/data/test | ./out/msh.local cat' 'StdinArn'

	@#./scripts/test.sh './out/msh.local echo "input"| ./out/sqs ' 'input'
	@#./scripts/test.sh 'cat test/data/lambda-input | ./out/msh.lambda sed "s/input/output/"' 'output'
	@./scripts/test.sh 'cat test/data/local-stdin | ./out/msh.pipes sed "s/input/output/"' 'test1'
#	@./scripts/test.sh 'echo "input"| ./out/sqs | ./out/msh.local sed "s/input/output/"' 'output'
	@#./scripts/test.sh 'echo "input"| ./out/sqs | ./out/msh.local bash ./test/cat.msh.sqs | ./out/sqs' 'output'
	@#./scripts/test.sh 'echo "input"| ./out/sqs | ./out/msh.local bash ./test/cat.msh.sqs | ./out/sqs' 'output'
	@#./scripts/test.sh 'echo "input"|./out/msh.lambda' 'output'
