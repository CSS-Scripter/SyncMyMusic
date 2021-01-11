package errors

import "app/config"

// BandDoesNotExistError is thrown when the ID given in a path is not a number
type BandDoesNotExistError struct{}

func (err *BandDoesNotExistError) Error() string {
	return config.ErrorMessages.BandDoesNotExist
}
