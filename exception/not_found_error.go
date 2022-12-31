package exception

type NotFoundError struct {
	Message string
}

func (notFoundError NotFoundError) Error() string {
	return notFoundError.Message
}
