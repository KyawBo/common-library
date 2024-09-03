package errors

type ForbiddenError struct {
	Message string `json:"message"`
}

func NewForbiddenError(err error) *ForbiddenError {
	return &ForbiddenError{
		Message: err.Error(),
	}
}

func (e *ForbiddenError) Error() string {
	return e.Message
}
