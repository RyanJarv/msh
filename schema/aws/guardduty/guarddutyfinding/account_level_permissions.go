package guarddutyfinding

type AccountLevelPermissions struct {
    BlockPublicAccess BlockPublicAccess `json:"blockPublicAccess,omitempty"`
}

func (a *AccountLevelPermissions) SetBlockPublicAccess(blockPublicAccess BlockPublicAccess) {
    a.BlockPublicAccess = blockPublicAccess
}
