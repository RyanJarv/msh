package voiceidspeakeraction

type ErrorInfo struct {
    ErrorCode float64 `json:"errorCode"`
    ErrorMessage string `json:"errorMessage"`
    ErrorType string `json:"errorType"`
}

func (e *ErrorInfo) SetErrorCode(errorCode float64) {
    e.ErrorCode = errorCode
}

func (e *ErrorInfo) SetErrorMessage(errorMessage string) {
    e.ErrorMessage = errorMessage
}

func (e *ErrorInfo) SetErrorType(errorType string) {
    e.ErrorType = errorType
}
