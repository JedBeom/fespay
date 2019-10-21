package models

type FieldError struct {
	message string
}

func (e FieldError) Error() string {
	return e.message
}
