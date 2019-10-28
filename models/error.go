package models

type FieldError struct {
	message string
}

func (e FieldError) Error() string {
	return e.message
}

func NewFieldError(m string) FieldError {
	return FieldError{m}
}
