package simpleworkflowexecutionstatechange

type TaskList struct {
    Name string `json:"name"`
}

func (t *TaskList) SetName(name string) {
    t.Name = name
}
