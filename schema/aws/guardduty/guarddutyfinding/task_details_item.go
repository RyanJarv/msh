package guarddutyfinding

type TaskDetailsItem struct {
    Image string `json:"image"`
    Name string `json:"name"`
}

func (t *TaskDetailsItem) SetImage(image string) {
    t.Image = image
}

func (t *TaskDetailsItem) SetName(name string) {
    t.Name = name
}
