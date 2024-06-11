package awsconsolesigninviacloudtrail

type ResponseElements struct {
    ConsoleLogin string `json:"ConsoleLogin"`
}

func (r *ResponseElements) SetConsoleLogin(consoleLogin string) {
    r.ConsoleLogin = consoleLogin
}
