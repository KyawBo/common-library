package errors

type FinalInternalServerError struct {
	Message string `json:"message"`
}

func NewFinalInternalServerError(err error) *FinalInternalServerError {
	return &FinalInternalServerError{
		Message: err.Error(),
	}
}

func (e *FinalInternalServerError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string `json:"message"`
}

func NewInternalServerError(err error) *InternalServerError {
	return &InternalServerError{
		Message: err.Error(),
	}
}

func (e *InternalServerError) Error() string {
	return e.Message
}
