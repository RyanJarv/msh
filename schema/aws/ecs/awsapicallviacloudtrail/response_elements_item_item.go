package awsapicallviacloudtrail

type ResponseElementsItemItem struct {
    Details []ResponseElementsItemItemItem `json:"details"`
    Id string `json:"id"`
    ResponseElementsItemItemType string `json:"type"`
    Status string `json:"status"`
}

func (r *ResponseElementsItemItem) SetDetails(details []ResponseElementsItemItemItem) {
    r.Details = details
}

func (r *ResponseElementsItemItem) SetId(id string) {
    r.Id = id
}

func (r *ResponseElementsItemItem) SetResponseElementsItemItemType(responseElementsItemItemType string) {
    r.ResponseElementsItemItemType = responseElementsItemItemType
}

func (r *ResponseElementsItemItem) SetStatus(status string) {
    r.Status = status
}
