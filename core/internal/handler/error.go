package handler

type Error interface {
	error
	Status() int
}

type HTTPError struct {
	Message error
	Code    int
}

func (h HTTPError) Error() string {
	return h.Message.Error()
}

func (h HTTPError) Status() int {
	return h.Code
}
