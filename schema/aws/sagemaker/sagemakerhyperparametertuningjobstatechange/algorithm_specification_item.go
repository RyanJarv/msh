package sagemakerhyperparametertuningjobstatechange

type AlgorithmSpecificationItem struct {
    Regex string `json:"Regex"`
    Name string `json:"Name"`
}

func (a *AlgorithmSpecificationItem) SetRegex(regex string) {
    a.Regex = regex
}

func (a *AlgorithmSpecificationItem) SetName(name string) {
    a.Name = name
}
