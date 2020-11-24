build:
	go build -o out/msh cmd/main.go

install: build
	cp out/msh /usr/local/bin/msh
	chmod +x /usr/local/bin/msh

test: build
	/Users/ryan/Code/msh/out/msh dockerfile /Users/ryan/Code/msh/config/echo.msh

echo: install
	DEBUG=true ./config/echo.msh hello world

cat: install
	DEBUG=true ./config/cat.msh
