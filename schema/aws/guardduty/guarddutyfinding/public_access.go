package guarddutyfinding

type PublicAccess struct {
    PermissionConfiguration PermissionConfiguration `json:"permissionConfiguration,omitempty"`
    EffectivePermission string `json:"effectivePermission,omitempty"`
}

func (p *PublicAccess) SetPermissionConfiguration(permissionConfiguration PermissionConfiguration) {
    p.PermissionConfiguration = permissionConfiguration
}

func (p *PublicAccess) SetEffectivePermission(effectivePermission string) {
    p.EffectivePermission = effectivePermission
}
