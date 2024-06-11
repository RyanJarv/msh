package cloudformationhookinvocationprogress

type Result struct {
    Data string `json:"data"`
}

func (r *Result) SetData(data string) {
    r.Data = data
}
