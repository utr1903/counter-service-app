package commons

// Result : For returning meaningful responses in case of failure
type CustomResult struct {
	Success bool        `json:"success"`
	Model   interface{} `json:"model"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Error   *error      `json:"error"`
}

// Success : General way to response with success
func Success(data interface{}, err error) *CustomResult {
	return &CustomResult{
		Success: true,
		Model:   data,
		Code:    "",
		Message: "",
		Error:   nil,
	}
}

// BadRequest : Request structure is not in desired format
func BadRequest(err *error) *CustomResult {
	return &CustomResult{
		Success: false,
		Model:   nil,
		Code:    "BadRequest",
		Message: "Request structure not valid!",
		Error:   err,
	}
}

// InternalServerError : An internal failure is occured
func InternalServerError(err *error) *CustomResult {
	return &CustomResult{
		Success: false,
		Model:   nil,
		Code:    "InternalServerError",
		Message: "Internal server error occured!",
		Error:   err,
	}
}
