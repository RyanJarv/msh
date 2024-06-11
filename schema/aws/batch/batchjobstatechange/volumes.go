package batchjobstatechange

type Volumes struct {
    Name string `json:"name,omitempty"`
    Host Host `json:"host,omitempty"`
}

func (v *Volumes) SetName(name string) {
    v.Name = name
}

func (v *Volumes) SetHost(host Host) {
    v.Host = host
}
