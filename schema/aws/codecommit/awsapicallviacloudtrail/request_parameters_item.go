package awsapicallviacloudtrail

type RequestParametersItem struct {
    FilePath string `json:"filePath"`
}

func (r *RequestParametersItem) SetFilePath(filePath string) {
    r.FilePath = filePath
}
