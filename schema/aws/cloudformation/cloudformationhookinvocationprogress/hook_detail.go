package cloudformationhookinvocationprogress

type Hook_detail struct {
    Description string `json:"description"`
    DocumentationUrl string `json:"documentation-url"`
    FailureMode string `json:"failure-mode"`
    HookTypeArn string `json:"hook-type-arn"`
    HookTypeName string `json:"hook-type-name"`
    HookVersion string `json:"hook-version"`
    SourceUrl string `json:"source-url"`
}

func (h *Hook_detail) SetDescription(description string) {
    h.Description = description
}

func (h *Hook_detail) SetDocumentationUrl(documentationUrl string) {
    h.DocumentationUrl = documentationUrl
}

func (h *Hook_detail) SetFailureMode(failureMode string) {
    h.FailureMode = failureMode
}

func (h *Hook_detail) SetHookTypeArn(hookTypeArn string) {
    h.HookTypeArn = hookTypeArn
}

func (h *Hook_detail) SetHookTypeName(hookTypeName string) {
    h.HookTypeName = hookTypeName
}

func (h *Hook_detail) SetHookVersion(hookVersion string) {
    h.HookVersion = hookVersion
}

func (h *Hook_detail) SetSourceUrl(sourceUrl string) {
    h.SourceUrl = sourceUrl
}
