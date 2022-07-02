build:
	go generate ./...
	go build -o out/msh main.go

install: build
	mkdir -p ~/.msh/bin
	echo 'export PATH=$$HOME/.msh/bin/:$$PATH' > ~/.msh/env
	cp out/msh ~/.msh/bin/msh
	cp test/* ~/.msh/bin/
	chmod +x ~/.msh/bin/msh

test: build
	./scripts/test.sh './test/python.msh --version' 'Python 3.*'