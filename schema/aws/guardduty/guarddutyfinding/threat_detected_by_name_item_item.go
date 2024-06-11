package guarddutyfinding

type ThreatDetectedByNameItemItem struct {
    FileName string `json:"fileName"`
    FilePath string `json:"filePath"`
    Hash string `json:"hash"`
    VolumeArn string `json:"volumeArn"`
}

func (t *ThreatDetectedByNameItemItem) SetFileName(fileName string) {
    t.FileName = fileName
}

func (t *ThreatDetectedByNameItemItem) SetFilePath(filePath string) {
    t.FilePath = filePath
}

func (t *ThreatDetectedByNameItemItem) SetHash(hash string) {
    t.Hash = hash
}

func (t *ThreatDetectedByNameItemItem) SetVolumeArn(volumeArn string) {
    t.VolumeArn = volumeArn
}
