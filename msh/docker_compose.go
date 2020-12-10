package msh

import "log"

func Compose(argv []string) {
	log.Println("args: ", argv)
	cmdStr := append([]string{"docker-compose", "-f", "{{.Path}}", "run", "--service-ports", "app"}, argv[1:]...)
	Exec(argv[0], cmdStr)
}
