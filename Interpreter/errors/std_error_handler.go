package errors

type PanicErrorHandler struct {
}

func (peh PanicErrorHandler) ThrowError(message string) {
	panic(message)
}

func CreatePanicErrorHandler() PanicErrorHandler {
	return PanicErrorHandler{}
}

func (sapeh PanicErrorHandler) ThrowErrorIfFalse(message string, assertion bool) {
	panic(message)
}
