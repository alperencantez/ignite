package pkg

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type IgniteError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}
