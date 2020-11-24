build:
	go build -o out/msh cmd/main.go

install: build
	cp out/msh /usr/local/bin/msh
	chmod +x /usr/local/bin/msh

test: build
	/Users/ryan/Code/msh/out/msh dockerfile /Users/ryan/Code/msh/config/echo.msh

echo: build
	./config/echo.msh

cat: build
	./config/echo.msh
