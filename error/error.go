package error

type Error struct {
	status  int
	message string
	errors  any
}

func NewError(status int, message string, errors any) *Error {
	return &Error{
		status:  status,
		message: message,
		errors:  errors,
	}
}

func (e *Error) GetStatus() int {
	return e.status
}

func (e *Error) GetMessage() string {
	return e.message
}

func (e *Error) GetErrors() any {
	return e.errors
}
