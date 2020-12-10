package msh

func Compose(argv []string) {
	cmdStr := append([]string{"docker-compose", "-f", "{{.Path}}", "run", "app", "--"}, argv[1:]...)
	Exec(argv[0], cmdStr)
}
