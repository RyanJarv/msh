package voiceidbatchspeakerenrollmentaction

type ErrorInfo struct {
    ErrorCode float64 `json:"errorCode,omitempty"`
    ErrorMessage string `json:"errorMessage,omitempty"`
    ErrorType string `json:"errorType,omitempty"`
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
