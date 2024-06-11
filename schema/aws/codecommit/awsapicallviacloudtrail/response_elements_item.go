package awsapicallviacloudtrail

type ResponseElementsItem struct {
    AbsolutePath string `json:"absolutePath"`
    BlobId string `json:"blobId"`
    FileMode string `json:"fileMode"`
}

func (r *ResponseElementsItem) SetAbsolutePath(absolutePath string) {
    r.AbsolutePath = absolutePath
}

func (r *ResponseElementsItem) SetBlobId(blobId string) {
    r.BlobId = blobId
}

func (r *ResponseElementsItem) SetFileMode(fileMode string) {
    r.FileMode = fileMode
}
