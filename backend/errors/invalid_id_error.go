package errors

import "app/config"

// InvalidIDError is thrown when an invalid ID occurs
type InvalidIDError struct{}

func (err *InvalidIDError) Error() string {
	return config.ErrorMessages.InvalidID
}
