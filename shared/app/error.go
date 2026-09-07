package app

const (
	CodeNotFound      = "NOT_FOUND"
	CodeInvalidInput  = "INVALID_INPUT"
	CodeConflict      = "CONFLICT"
	CodeUnauthorized  = "UNAUTHORIZED"
	CodeForbidden     = "FORBIDDEN"
	CodeInternalError = "INTERNAL_ERROR"
)

type AppError struct {
	Message string
	Code    string
	Details []string

	Err error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}
