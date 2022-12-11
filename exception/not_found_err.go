package exception

type NotFoundError struct {
	Err error
}

func NewNotFound(err error) NotFoundError {
	return NotFoundError{Err: err}
}
