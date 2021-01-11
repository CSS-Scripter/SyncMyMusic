package errors

import "app/config"

// IDIsNotANumberError is thrown when the ID given in a path is not a number
type IDIsNotANumberError struct{}

func (err *IDIsNotANumberError) Error() string {
	return config.ErrorMessages.IDIsNotANumber
}
