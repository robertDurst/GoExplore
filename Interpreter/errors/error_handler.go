package errors

type ErrorHandler interface {
	ThrowError(message string)
}
