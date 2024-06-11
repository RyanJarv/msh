package chimevoiceconnectorstreamingstatus

import (
    "time"
)


type ChimeVoiceConnectorStreamingStatus struct {
    CallId string `json:"callId"`
    Direction string `json:"direction"`
    MediaType string `json:"mediaType"`
    StartFragmentNumber string `json:"startFragmentNumber"`
    StartTime time.Time `json:"startTime"`
    StreamArn string `json:"streamArn"`
    StreamingStatus string `json:"streamingStatus"`
    TransactionId string `json:"transactionId"`
    Version string `json:"version"`
    VoiceConnectorId string `json:"voiceConnectorId"`
    InviteHeaders map[string]string `json:"inviteHeaders,omitempty"`
    SiprecMetadata string `json:"siprecMetadata,omitempty"`
}

func (c *ChimeVoiceConnectorStreamingStatus) SetCallId(callId string) {
    c.CallId = callId
}

func (c *ChimeVoiceConnectorStreamingStatus) SetDirection(direction string) {
    c.Direction = direction
}

func (c *ChimeVoiceConnectorStreamingStatus) SetMediaType(mediaType string) {
    c.MediaType = mediaType
}

func (c *ChimeVoiceConnectorStreamingStatus) SetStartFragmentNumber(startFragmentNumber string) {
    c.StartFragmentNumber = startFragmentNumber
}

func (c *ChimeVoiceConnectorStreamingStatus) SetStartTime(startTime time.Time) {
    c.StartTime = startTime
}

func (c *ChimeVoiceConnectorStreamingStatus) SetStreamArn(streamArn string) {
    c.StreamArn = streamArn
}

func (c *ChimeVoiceConnectorStreamingStatus) SetStreamingStatus(streamingStatus string) {
    c.StreamingStatus = streamingStatus
}

func (c *ChimeVoiceConnectorStreamingStatus) SetTransactionId(transactionId string) {
    c.TransactionId = transactionId
}

func (c *ChimeVoiceConnectorStreamingStatus) SetVersion(version string) {
    c.Version = version
}

func (c *ChimeVoiceConnectorStreamingStatus) SetVoiceConnectorId(voiceConnectorId string) {
    c.VoiceConnectorId = voiceConnectorId
}

func (c *ChimeVoiceConnectorStreamingStatus) SetInviteHeaders(inviteHeaders map[string]string) {
    c.InviteHeaders = inviteHeaders
}

func (c *ChimeVoiceConnectorStreamingStatus) SetSiprecMetadata(siprecMetadata string) {
    c.SiprecMetadata = siprecMetadata
}
