package errors

type ErrorHandler interface {
	ThrowError(message string)
	ThrowErrorIfFalse(message string, shouldThrow bool)
}
