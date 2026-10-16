package app

type AppResponse[T any] struct {
	Success bool       `json:"success"`
	Data    *T         `json:"data,omitempty"`
	Meta    any        `json:"meta,omitempty"`
	Error   *ErrorInfo `json:"error,omitempty"`

	// for development
	Stack *string `json:"stack,omitempty"`
}

type ErrorInfo struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Details []ValidationError `json:"details,omitempty"`
}

type ValidationError struct {
	Field   string   `json:"field"`
	Message []string `json:"message"`
}
