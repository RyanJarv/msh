package guarddutyfinding

type NewPolicy struct {
    AllowUsersToChangePassword string `json:"allowUsersToChangePassword,omitempty"`
    HardExpiry string `json:"hardExpiry,omitempty"`
    MaxPasswordAge string `json:"maxPasswordAge,omitempty"`
    MinimumPasswordLength string `json:"minimumPasswordLength,omitempty"`
    PasswordReusePrevention string `json:"passwordReusePrevention,omitempty"`
    RequireLowercaseCharacters string `json:"requireLowercaseCharacters,omitempty"`
    RequireNumbers string `json:"requireNumbers,omitempty"`
    RequireSymbols string `json:"requireSymbols,omitempty"`
    RequireUppercaseCharacters string `json:"requireUppercaseCharacters,omitempty"`
}

func (n *NewPolicy) SetAllowUsersToChangePassword(allowUsersToChangePassword string) {
    n.AllowUsersToChangePassword = allowUsersToChangePassword
}

func (n *NewPolicy) SetHardExpiry(hardExpiry string) {
    n.HardExpiry = hardExpiry
}

func (n *NewPolicy) SetMaxPasswordAge(maxPasswordAge string) {
    n.MaxPasswordAge = maxPasswordAge
}

func (n *NewPolicy) SetMinimumPasswordLength(minimumPasswordLength string) {
    n.MinimumPasswordLength = minimumPasswordLength
}

func (n *NewPolicy) SetPasswordReusePrevention(passwordReusePrevention string) {
    n.PasswordReusePrevention = passwordReusePrevention
}

func (n *NewPolicy) SetRequireLowercaseCharacters(requireLowercaseCharacters string) {
    n.RequireLowercaseCharacters = requireLowercaseCharacters
}

func (n *NewPolicy) SetRequireNumbers(requireNumbers string) {
    n.RequireNumbers = requireNumbers
}

func (n *NewPolicy) SetRequireSymbols(requireSymbols string) {
    n.RequireSymbols = requireSymbols
}

func (n *NewPolicy) SetRequireUppercaseCharacters(requireUppercaseCharacters string) {
    n.RequireUppercaseCharacters = requireUppercaseCharacters
}
