package awsapicallviacloudtrail

type Location struct {
    FilePath string `json:"filePath,omitempty"`
    FilePosition float64 `json:"filePosition,omitempty"`
    RelativeFileVersion string `json:"relativeFileVersion,omitempty"`
}

func (l *Location) SetFilePath(filePath string) {
    l.FilePath = filePath
}

func (l *Location) SetFilePosition(filePosition float64) {
    l.FilePosition = filePosition
}

func (l *Location) SetRelativeFileVersion(relativeFileVersion string) {
    l.RelativeFileVersion = relativeFileVersion
}
