package awsapicallviacloudtrail

type ResponseElements struct {
    InstancesSet InstancesSet `json:"instancesSet,omitempty"`
    CreateFleetResponse CreateFleetResponse `json:"CreateFleetResponse,omitempty"`
    DeleteLaunchTemplateResponse DeleteLaunchTemplateResponse `json:"DeleteLaunchTemplateResponse,omitempty"`
    NetworkInterface NetworkInterface `json:"networkInterface,omitempty"`
    CreateLaunchTemplateResponse DeleteLaunchTemplateResponse `json:"CreateLaunchTemplateResponse,omitempty"`
    GroupSet interface{} `json:"groupSet,omitempty"`
    ResponseElementsReturn bool `json:"_return,omitempty"`
    RequesterId string `json:"requesterId,omitempty"`
    ReservationId string `json:"reservationId,omitempty"`
    RequestId string `json:"requestId,omitempty"`
    GroupId string `json:"groupId,omitempty"`
    OwnerId string `json:"ownerId,omitempty"`
}

func (r *ResponseElements) SetInstancesSet(instancesSet InstancesSet) {
    r.InstancesSet = instancesSet
}

func (r *ResponseElements) SetCreateFleetResponse(createFleetResponse CreateFleetResponse) {
    r.CreateFleetResponse = createFleetResponse
}

func (r *ResponseElements) SetDeleteLaunchTemplateResponse(deleteLaunchTemplateResponse DeleteLaunchTemplateResponse) {
    r.DeleteLaunchTemplateResponse = deleteLaunchTemplateResponse
}

func (r *ResponseElements) SetNetworkInterface(networkInterface NetworkInterface) {
    r.NetworkInterface = networkInterface
}

func (r *ResponseElements) SetCreateLaunchTemplateResponse(createLaunchTemplateResponse DeleteLaunchTemplateResponse) {
    r.CreateLaunchTemplateResponse = createLaunchTemplateResponse
}

func (r *ResponseElements) SetGroupSet(groupSet interface{}) {
    r.GroupSet = groupSet
}

func (r *ResponseElements) SetResponseElementsReturn(responseElementsReturn bool) {
    r.ResponseElementsReturn = responseElementsReturn
}

func (r *ResponseElements) SetRequesterId(requesterId string) {
    r.RequesterId = requesterId
}

func (r *ResponseElements) SetReservationId(reservationId string) {
    r.ReservationId = reservationId
}

func (r *ResponseElements) SetRequestId(requestId string) {
    r.RequestId = requestId
}

func (r *ResponseElements) SetGroupId(groupId string) {
    r.GroupId = groupId
}

func (r *ResponseElements) SetOwnerId(ownerId string) {
    r.OwnerId = ownerId
}
