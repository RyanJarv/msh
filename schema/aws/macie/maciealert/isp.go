package maciealert

type ISP struct {
    Amazon float64 `json:"Amazon,omitempty"`
}

func (i *ISP) SetAmazon(amazon float64) {
    i.Amazon = amazon
}
