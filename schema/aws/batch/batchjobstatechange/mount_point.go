package batchjobstatechange

type MountPoint struct {
    ContainerPath string `json:"containerPath,omitempty"`
    ReadOnly bool `json:"readOnly,omitempty"`
    SourceVolume string `json:"sourceVolume,omitempty"`
}

func (m *MountPoint) SetContainerPath(containerPath string) {
    m.ContainerPath = containerPath
}

func (m *MountPoint) SetReadOnly(readOnly bool) {
    m.ReadOnly = readOnly
}

func (m *MountPoint) SetSourceVolume(sourceVolume string) {
    m.SourceVolume = sourceVolume
}
