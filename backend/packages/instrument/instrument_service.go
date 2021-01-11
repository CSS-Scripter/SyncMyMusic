package instrument

import (
	dao "app/packages/instrument/dao"
	"app/structs"
)

// CreateInstrument is used to create a Musician (User) for the application.
func CreateInstrument(instrument *structs.Instrument) error {
	return dao.CreateInstrument(instrument)
}

// DeleteInstrument deletes an instrument.
func DeleteInstrument(instrumentID int) error {
	return dao.DeleteInstrument(instrumentID)
}

// GetInstrumentsForBand is used to get all Instruments for the band.
func GetInstrumentsForBand(bandID int) ([]structs.Instrument, error) {
	return dao.GetInstrumentsForBand(bandID)
}
