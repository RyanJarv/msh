package guarddutyfinding

type EbsVolumeDetailsItem struct {
    DeviceName string `json:"deviceName"`
    EncryptionType string `json:"encryptionType"`
    KmsKeyArn interface{} `json:"kmsKeyArn"`
    SnapshotArn string `json:"snapshotArn"`
    VolumeArn string `json:"volumeArn"`
    VolumeSizeInGB float64 `json:"volumeSizeInGB"`
    VolumeType string `json:"volumeType"`
}

func (e *EbsVolumeDetailsItem) SetDeviceName(deviceName string) {
    e.DeviceName = deviceName
}

func (e *EbsVolumeDetailsItem) SetEncryptionType(encryptionType string) {
    e.EncryptionType = encryptionType
}

func (e *EbsVolumeDetailsItem) SetKmsKeyArn(kmsKeyArn interface{}) {
    e.KmsKeyArn = kmsKeyArn
}

func (e *EbsVolumeDetailsItem) SetSnapshotArn(snapshotArn string) {
    e.SnapshotArn = snapshotArn
}

func (e *EbsVolumeDetailsItem) SetVolumeArn(volumeArn string) {
    e.VolumeArn = volumeArn
}

func (e *EbsVolumeDetailsItem) SetVolumeSizeInGB(volumeSizeInGB float64) {
    e.VolumeSizeInGB = volumeSizeInGB
}

func (e *EbsVolumeDetailsItem) SetVolumeType(volumeType string) {
    e.VolumeType = volumeType
}
