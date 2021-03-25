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

// --- Not Valid Exceptions ---

// RequestNotValid : Request structure is not in desired format
func RequestNotValid() *CustomResult {
	return &CustomResult{
		Success: false,
		Model:   nil,
		Code:    "RequestNotValid",
		Message: "Request structure not valid!",
		Error:   nil,
	}
}

// ----------------------------
