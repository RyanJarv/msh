package awsapicallviacloudtrail

type RequestParametersItem_2 struct {
    DataSource DataSource_1 `json:"dataSource"`
    ChannelName string `json:"channelName"`
    ContentType string `json:"contentType"`
}

func (r *RequestParametersItem_2) SetDataSource(dataSource DataSource_1) {
    r.DataSource = dataSource
}

func (r *RequestParametersItem_2) SetChannelName(channelName string) {
    r.ChannelName = channelName
}

func (r *RequestParametersItem_2) SetContentType(contentType string) {
    r.ContentType = contentType
}
