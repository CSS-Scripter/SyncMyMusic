package errors

import "app/config"

// NoDatabaseEntryError is thrown when there is no entry found on a crucial moment
type NoDatabaseEntryError struct{}

func (err *NoDatabaseEntryError) Error() string {
	return config.ErrorMessages.NoDatabaseEntry
}
