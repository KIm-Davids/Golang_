package exceptions

type InvalidDetailsException struct {
	Message string
}

func (e *InvalidDetailsException) Error() string {
	return e.Message
}
