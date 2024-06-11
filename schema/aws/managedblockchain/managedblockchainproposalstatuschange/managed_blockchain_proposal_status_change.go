package managedblockchainproposalstatuschange

import (
    "time"
)


type ManagedBlockchainProposalStatusChange struct {
    ProposedByMemberName string `json:"proposedByMemberName"`
    ProposedByMemberId string `json:"proposedByMemberId"`
    Description string `json:"description,omitempty"`
    NetworkId string `json:"networkId"`
    Message string `json:"message"`
    ProposalId string `json:"proposalId"`
    ExpirationDate time.Time `json:"expirationDate"`
    Status string `json:"status"`
}

func (m *ManagedBlockchainProposalStatusChange) SetProposedByMemberName(proposedByMemberName string) {
    m.ProposedByMemberName = proposedByMemberName
}

func (m *ManagedBlockchainProposalStatusChange) SetProposedByMemberId(proposedByMemberId string) {
    m.ProposedByMemberId = proposedByMemberId
}

func (m *ManagedBlockchainProposalStatusChange) SetDescription(description string) {
    m.Description = description
}

func (m *ManagedBlockchainProposalStatusChange) SetNetworkId(networkId string) {
    m.NetworkId = networkId
}

func (m *ManagedBlockchainProposalStatusChange) SetMessage(message string) {
    m.Message = message
}

func (m *ManagedBlockchainProposalStatusChange) SetProposalId(proposalId string) {
    m.ProposalId = proposalId
}

func (m *ManagedBlockchainProposalStatusChange) SetExpirationDate(expirationDate time.Time) {
    m.ExpirationDate = expirationDate
}

func (m *ManagedBlockchainProposalStatusChange) SetStatus(status string) {
    m.Status = status
}
