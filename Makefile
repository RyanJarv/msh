export MSH_SKIPDEPLOY=1

.PHONY: build clean

pull:
	git submodule update --init

build: clean
	go generate ./...
	mkdir -p out
	go build -o out/msh cmd/*.go
	./scripts/link_cmds.sh ./out

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

gen-test-input:
	./out/@cron '0 0 * * * *' > ./test/inputs/cron.json
	./out/@cron '0 0 * * * *' | ./out/.sfn > ./test/inputs/sfn.json
	./out/@cron '0 0 * * * *' | ./out/.sfn | ./out/.aws ec2 describe-instances --query 'Reservations[]. Instances[]' > ./test/inputs/aws.json
	./out/@cron '0 0 * * * *' | ./out/.sfn | ./out/.aws ec2 describe-instances --query 'Reservations[]. Instances[]' | ./out/.foreach > ./test/inputs/foreach.json
	./out/@cron '0 0 * * * *' | ./out/.sfn | ./out/.aws ec2 describe-instances --query 'Reservations[]. Instances[]' | ./out/.foreach | ./out/.filter ./scripts/hours_running.py 72 > ./test/inputs/filter.json

test: test-name test-sfn test-each test-sleep test-lambda
	@echo "\nRunning Tests\n"

test-name:
	@DEBUG=1 ./scripts/test.sh './out/.name asdf | ./out/.sleep 10 | cat' '"Name":"asdf"'

test-sfn:
	@DEBUG=1 ./scripts/test.sh 'cat ./test/inputs/cron.json | ./out/.sfn | cat' '"Name":"sfn"'

test-each:
	@DEBUG=1 ./scripts/test.sh 'cat ./test/inputs/aws.json | ./out/.foreach | cat' '"Name":"foreach"'

test-sleep:
	@DEBUG=1 ./scripts/test.sh 'cat ./test/inputs/sfn.json | ./out/.sleep 3 | cat' '"Name":"sleep"'

test-lambda:
	@DEBUG=1 ./scripts/test.sh 'cat ./test/inputs/sfn.json | ./out/.lambda ./scripts/hours_running.py 72 | cat' '"Name":"lambda"'

install: build
	./scripts/link_cmds.sh ~/.msh/bin
	test -f ~/.zshrc && if ! fgrep -q 'PATH="$$PATH:~/.msh/bin"' ~/.zshrc; then echo 'export PATH="$$PATH:~/.msh/bin"' >> ~/.zshrc; fi
	test -f ~/.bashrc && if ! fgrep -q 'PATH="$$PATH:~/.msh/bin"' ~/.zshrc; then echo 'export PATH="$$PATH:~/.msh/bin"' >> ~/.bashrc; fi
