package errors

type UnauthorizedError struct {
	Message string `json:"message"`
}

func NewUnauthorizedError(err error) *UnauthorizedError {
	return &UnauthorizedError{
		Message: err.Error(),
	}
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}
