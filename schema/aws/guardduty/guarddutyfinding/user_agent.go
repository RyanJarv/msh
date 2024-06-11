package guarddutyfinding

type UserAgent struct {
    FullUserAgent string `json:"fullUserAgent,omitempty"`
    UserAgentCategory string `json:"userAgentCategory,omitempty"`
}

func (u *UserAgent) SetFullUserAgent(fullUserAgent string) {
    u.FullUserAgent = fullUserAgent
}

func (u *UserAgent) SetUserAgentCategory(userAgentCategory string) {
    u.UserAgentCategory = userAgentCategory
}
