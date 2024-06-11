package awsapicallviacloudtrail

type RequestParameters struct {
    AuthCode string `json:"AuthCode,omitempty"`
    CodeVerifier string `json:"CodeVerifier,omitempty"`
    EnvironmentId string `json:"EnvironmentId"`
    FileDownloadPath string `json:"FileDownloadPath,omitempty"`
    FileUploadPath string `json:"FileUploadPath,omitempty"`
    KeyBase string `json:"KeyBase,omitempty"`
    RedirectUri string `json:"RedirectUri,omitempty"`
    RefreshToken string `json:"RefreshToken,omitempty"`
}

func (r *RequestParameters) SetAuthCode(authCode string) {
    r.AuthCode = authCode
}

func (r *RequestParameters) SetCodeVerifier(codeVerifier string) {
    r.CodeVerifier = codeVerifier
}

func (r *RequestParameters) SetEnvironmentId(environmentId string) {
    r.EnvironmentId = environmentId
}

func (r *RequestParameters) SetFileDownloadPath(fileDownloadPath string) {
    r.FileDownloadPath = fileDownloadPath
}

func (r *RequestParameters) SetFileUploadPath(fileUploadPath string) {
    r.FileUploadPath = fileUploadPath
}

func (r *RequestParameters) SetKeyBase(keyBase string) {
    r.KeyBase = keyBase
}

func (r *RequestParameters) SetRedirectUri(redirectUri string) {
    r.RedirectUri = redirectUri
}

func (r *RequestParameters) SetRefreshToken(refreshToken string) {
    r.RefreshToken = refreshToken
}
