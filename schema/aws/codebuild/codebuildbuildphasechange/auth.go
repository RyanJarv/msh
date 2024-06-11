package codebuildbuildphasechange

type Auth struct {
    AuthType string `json:"type,omitempty"`
}

func (a *Auth) SetAuthType(authType string) {
    a.AuthType = authType
}
