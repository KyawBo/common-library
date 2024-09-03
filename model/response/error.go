package response

type ErrorResponse struct {
	Status ErrorResponseStatus `json:"status"`
}

type ErrorResponseStatus struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type StandardErrorCode struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func BuildErrorResponse(standardErrorCode StandardErrorCode) ErrorResponse {
	return ErrorResponse{
		Status: ErrorResponseStatus(standardErrorCode),
	}
}
