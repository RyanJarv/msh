package codebuildbuildstatechange

type Auth struct {
    AuthType string `json:"type,omitempty"`
    Resource string `json:"resource,omitempty"`
}

func (a *Auth) SetAuthType(authType string) {
    a.AuthType = authType
}

func (a *Auth) SetResource(resource string) {
    a.Resource = resource
}
