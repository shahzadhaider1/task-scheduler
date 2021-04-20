package errors

// const will be used for the status of api handlers
const (
	NotFound            = "NOT_FOUND"
	InternalServerError = "INTERNAL_SERVER_ERROR"
)

// APIError struct contains the code and message of error
type APIError struct {
	code    string
	message string
}

func (a *APIError) Error() string {
	return a.message
}

// IsError will return whether error exists or not
func (a *APIError) IsError(errType string) bool {
	return a.code == errType
}

// NewAPIError returns the error type and error message
func NewAPIError(errType string, message string) *APIError {
	return &APIError{
		code:    errType,
		message: message,
	}
}
