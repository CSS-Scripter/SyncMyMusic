package instrumentdao

import (
	"app/database"
	"app/structs"
)

// CreateInstrument is used to create an instrument for the application.
func CreateInstrument(instrument *structs.Instrument) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	_, err := db.Model(instrument).Insert()
	return err
}

// DeleteInstrument deletes an instrument from a band.
func DeleteInstrument(instrumentID int) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	var err error
	_, err = db.Model(&structs.Instrument{}).Where("id = ?", instrumentID).Delete()
	return err
}
