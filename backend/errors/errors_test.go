package errors

import (
	"app/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IDIsNotANumberError(t *testing.T) {
	stackTrace := new(IDIsNotANumberError).Error()
	assert.Equal(t, stackTrace, config.ErrorMessages.IDIsNotANumber)
}

func Test_NoDatabaseEntryError(t *testing.T) {
	stackTrace := new(NoDatabaseEntryError).Error()
	assert.Equal(t, stackTrace, config.ErrorMessages.NoDatabaseEntry)
}

func Test_InvalidIDError(t *testing.T) {
	stackTrace := new(InvalidIDError).Error()
	assert.Equal(t, stackTrace, config.ErrorMessages.InvalidID)
}

func Test_BandDoesNotExistError(t *testing.T) {
	stackTrace := new(BandDoesNotExistError).Error()
	assert.Equal(t, stackTrace, config.ErrorMessages.BandDoesNotExist)
}
