package response

import (
	"time"
)

type ErrorBody struct {
	ErrorCode string   `json:"errorCode"`
	Message   string   `json:"message"`
	Details   []string `json:"details,omitempty"`
	Stack     string   `json:"stack,omitempty"`
}

type ErrorResponse struct {
	Status    bool      `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Error     ErrorBody `json:"error"`
}
