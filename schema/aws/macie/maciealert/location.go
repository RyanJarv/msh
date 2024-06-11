package maciealert

type Location struct {
    UsEast1 float64 `json:"us-east-1,omitempty"`
}

func (l *Location) SetUsEast1(usEast1 float64) {
    l.UsEast1 = usEast1
}
