package errors

import "fmt"

type StdErrorHandler struct {
}

func (seh StdErrorHandler) ThrowError(message string) {
	fmt.Printf("[ERROR]: %s\n", message)
}

func CreateStdErrorHandler() StdErrorHandler {
	return StdErrorHandler{}
}
