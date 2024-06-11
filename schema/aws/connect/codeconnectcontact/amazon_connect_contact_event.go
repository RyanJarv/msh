package codeconnectcontact

type AmazonConnectContactEvent struct {
    AgentInfo AgentInfo `json:"agentInfo"`
    QueueInfo QueueInfo `json:"queueInfo"`
    Channel string `json:"channel"`
    ContactId string `json:"contactId"`
    EventType string `json:"eventType"`
    InitialContactId string `json:"initialContactId"`
    InitiationMethod string `json:"initiationMethod"`
    InstanceArn string `json:"instanceArn"`
    PreviousContactId string `json:"previousContactId"`
}

func (a *AmazonConnectContactEvent) SetAgentInfo(agentInfo AgentInfo) {
    a.AgentInfo = agentInfo
}

func (a *AmazonConnectContactEvent) SetQueueInfo(queueInfo QueueInfo) {
    a.QueueInfo = queueInfo
}

func (a *AmazonConnectContactEvent) SetChannel(channel string) {
    a.Channel = channel
}

func (a *AmazonConnectContactEvent) SetContactId(contactId string) {
    a.ContactId = contactId
}

func (a *AmazonConnectContactEvent) SetEventType(eventType string) {
    a.EventType = eventType
}

func (a *AmazonConnectContactEvent) SetInitialContactId(initialContactId string) {
    a.InitialContactId = initialContactId
}

func (a *AmazonConnectContactEvent) SetInitiationMethod(initiationMethod string) {
    a.InitiationMethod = initiationMethod
}

func (a *AmazonConnectContactEvent) SetInstanceArn(instanceArn string) {
    a.InstanceArn = instanceArn
}

func (a *AmazonConnectContactEvent) SetPreviousContactId(previousContactId string) {
    a.PreviousContactId = previousContactId
}
