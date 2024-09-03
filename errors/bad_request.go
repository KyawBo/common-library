package errors

type BadRequestError struct {
	Message string `json:"message"`
}

func NewBadRequestError(err error) *BadRequestError {
	return &BadRequestError{
		Message: err.Error(),
	}
}

func (e *BadRequestError) Error() string {
	return e.Message
}
