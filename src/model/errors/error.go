package errors

// Error failed info in system
type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string {
	return e.Message
}

// NewStringError new a error with message
func NewStringError(msg string, code int) error {
	return Error{
		Message: msg,
		Code:    code,
	}
}
