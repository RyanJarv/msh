package awsapicallviacloudtrail

type Media struct {
    MediaFileUri string `json:"mediaFileUri,omitempty"`
}

func (m *Media) SetMediaFileUri(mediaFileUri string) {
    m.MediaFileUri = mediaFileUri
}
