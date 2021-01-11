package instrumentdao

import (
	"app/database"
	"app/structs"
)

// GetInstrumentsForBand is used to get all instruments for a given band.
func GetInstrumentsForBand(bandID int) ([]structs.Instrument, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	result := []structs.Instrument{}
	err := db.Model(&result).Where("band_id = ?", bandID).Select()

	return result, err
}
