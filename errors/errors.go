package errors

type Error struct {
	status  int
	message string
	details any
}

func NewError(status int, message string, details any) *Error {
	return &Error{
		status:  status,
		message: message,
		details: details,
	}
}

func (e *Error) GetStatus() int {
	return e.status
}

func (e *Error) GetMessage() string {
	return e.message
}

func (e *Error) GetDetails() any {
	return e.details
}
