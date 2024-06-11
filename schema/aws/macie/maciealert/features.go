package maciealert

type Features struct {
    DistinctEventName DistinctEventName `json:"distinctEventName,omitempty"`
    ListInstanceProfilesForRole ListInstanceProfilesForRole `json:"ListInstanceProfilesForRole,omitempty"`
    ListRolePolicies ListInstanceProfilesForRole `json:"ListRolePolicies,omitempty"`
}

func (f *Features) SetDistinctEventName(distinctEventName DistinctEventName) {
    f.DistinctEventName = distinctEventName
}

func (f *Features) SetListInstanceProfilesForRole(listInstanceProfilesForRole ListInstanceProfilesForRole) {
    f.ListInstanceProfilesForRole = listInstanceProfilesForRole
}

func (f *Features) SetListRolePolicies(listRolePolicies ListInstanceProfilesForRole) {
    f.ListRolePolicies = listRolePolicies
}
