package errors

type RequestTimeoutError struct {
	Message string `json:"message"`
}

func NewRequestTimeoutError(err error) *RequestTimeoutError {
	return &RequestTimeoutError{
		Message: err.Error(),
	}
}

func (e *RequestTimeoutError) Error() string {
	return e.Message
}
