package guarddutyfinding

type OldPolicy struct {
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

func (o *OldPolicy) SetAllowUsersToChangePassword(allowUsersToChangePassword string) {
    o.AllowUsersToChangePassword = allowUsersToChangePassword
}

func (o *OldPolicy) SetHardExpiry(hardExpiry string) {
    o.HardExpiry = hardExpiry
}

func (o *OldPolicy) SetMaxPasswordAge(maxPasswordAge string) {
    o.MaxPasswordAge = maxPasswordAge
}

func (o *OldPolicy) SetMinimumPasswordLength(minimumPasswordLength string) {
    o.MinimumPasswordLength = minimumPasswordLength
}

func (o *OldPolicy) SetPasswordReusePrevention(passwordReusePrevention string) {
    o.PasswordReusePrevention = passwordReusePrevention
}

func (o *OldPolicy) SetRequireLowercaseCharacters(requireLowercaseCharacters string) {
    o.RequireLowercaseCharacters = requireLowercaseCharacters
}

func (o *OldPolicy) SetRequireNumbers(requireNumbers string) {
    o.RequireNumbers = requireNumbers
}

func (o *OldPolicy) SetRequireSymbols(requireSymbols string) {
    o.RequireSymbols = requireSymbols
}

func (o *OldPolicy) SetRequireUppercaseCharacters(requireUppercaseCharacters string) {
    o.RequireUppercaseCharacters = requireUppercaseCharacters
}
