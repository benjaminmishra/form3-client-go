package f3client

import (
	"fmt"
)

// ArgumentError is raised when the called of a function in this library misses
// to pass an mandatory argument or misses to pass a mandatory field in a request
type ArgumentError struct {
	arg     string
	message string
}

// NewArgError creates a new ArgumentError object
// It acceots the missing argument name and the validation message that needs to
// be passed back alog with the error
func NewArgError(missingArg string, validationMsg string) *ArgumentError {

	return &ArgumentError{
		arg:     missingArg,
		message: validationMsg,
	}
}

func (ae *ArgumentError) Error() string {
	return fmt.Sprintf("%s : %s", ae.arg, ae.message)
}
