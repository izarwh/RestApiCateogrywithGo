package exception

type InternalServerError struct {
	Err error
}

func NewInternalServerErr(err error) InternalServerError {
	return InternalServerError{Err: err}
}
