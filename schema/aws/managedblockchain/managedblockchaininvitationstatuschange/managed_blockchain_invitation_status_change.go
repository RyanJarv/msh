package managedblockchaininvitationstatuschange

import (
    "time"
)


type ManagedBlockchainInvitationStatusChange struct {
    CreatedAt time.Time `json:"createdAt,omitempty"`
    NetworkId string `json:"networkId"`
    InvitationId string `json:"invitationId"`
    Message string `json:"message,omitempty"`
    ExpiresAt time.Time `json:"expiresAt"`
    AcceptedAt time.Time `json:"acceptedAt,omitempty"`
    RejectedAt time.Time `json:"rejectedAt,omitempty"`
    Status string `json:"status"`
}

func (m *ManagedBlockchainInvitationStatusChange) SetCreatedAt(createdAt time.Time) {
    m.CreatedAt = createdAt
}

func (m *ManagedBlockchainInvitationStatusChange) SetNetworkId(networkId string) {
    m.NetworkId = networkId
}

func (m *ManagedBlockchainInvitationStatusChange) SetInvitationId(invitationId string) {
    m.InvitationId = invitationId
}

func (m *ManagedBlockchainInvitationStatusChange) SetMessage(message string) {
    m.Message = message
}

func (m *ManagedBlockchainInvitationStatusChange) SetExpiresAt(expiresAt time.Time) {
    m.ExpiresAt = expiresAt
}

func (m *ManagedBlockchainInvitationStatusChange) SetAcceptedAt(acceptedAt time.Time) {
    m.AcceptedAt = acceptedAt
}

func (m *ManagedBlockchainInvitationStatusChange) SetRejectedAt(rejectedAt time.Time) {
    m.RejectedAt = rejectedAt
}

func (m *ManagedBlockchainInvitationStatusChange) SetStatus(status string) {
    m.Status = status
}
