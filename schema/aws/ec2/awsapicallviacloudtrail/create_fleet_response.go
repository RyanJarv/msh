package awsapicallviacloudtrail

type CreateFleetResponse struct {
    FleetInstanceSet interface{} `json:"fleetInstanceSet,omitempty"`
    Xmlns string `json:"xmlns,omitempty"`
    RequestId string `json:"requestId,omitempty"`
    FleetId string `json:"fleetId,omitempty"`
    ErrorSet interface{} `json:"errorSet,omitempty"`
}

func (c *CreateFleetResponse) SetFleetInstanceSet(fleetInstanceSet interface{}) {
    c.FleetInstanceSet = fleetInstanceSet
}

func (c *CreateFleetResponse) SetXmlns(xmlns string) {
    c.Xmlns = xmlns
}

func (c *CreateFleetResponse) SetRequestId(requestId string) {
    c.RequestId = requestId
}

func (c *CreateFleetResponse) SetFleetId(fleetId string) {
    c.FleetId = fleetId
}

func (c *CreateFleetResponse) SetErrorSet(errorSet interface{}) {
    c.ErrorSet = errorSet
}
