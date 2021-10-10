package f3client

import "fmt"

type ArgumentError struct {
	arg     string
	message string
}

func NewArgError(missingArg string, validationMsg string) *ArgumentError {

	return &ArgumentError{
		arg:     missingArg,
		message: validationMsg,
	}
}

func (ae *ArgumentError) Error() string {
	return fmt.Sprintf("%s : %s", ae.arg, ae.message)
}
