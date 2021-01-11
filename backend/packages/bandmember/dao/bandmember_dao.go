package bandmemberdao

import (
	"app/database"
	"app/structs"
)

// GetBandmembers returns all bandmembers which are members of band with bandID
func GetBandmembers(bandID int) ([]structs.BandMember, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	members := []structs.BandMember{}
	query := db.Model(&members)
	query.ColumnExpr("instruments.instrument_name AS instrument__instrument_name, instruments.id AS instrument__id, instruments.band_id AS instrument__band_id")
	query.ColumnExpr("musicians.id AS musician__id, musicians.musician_name AS musician__musician_name, musicians.email AS musician__email")
	query.ColumnExpr("bands.id AS band__id, bands.band_name AS band__band_name")
	query.Join("LEFT JOIN instruments").JoinOn("bandmember.instrument_id = instruments.id")
	query.Join("LEFT JOIN musicians").JoinOn("bandmember.musician_id = musicians.id")
	query.Join("LEFT JOIN bands").JoinOn("bandmember.band_id = bands.id")
	query.Where("bandmember.band_id = ?", bandID)
	err := query.Select()

	return members, err
}

// UpdateBandmember updated the bandmember in the db with the new data
func UpdateBandmember(bandmember structs.BandMember) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	_, err := db.Model(&bandmember).Where("bandmember.band_id = ?", bandmember.BandID).Where("bandmember.musician_id = ?", bandmember.MusicianID).Update()

	return err
}

// DeleteBandmember used to delete a single member based on bandid and musicianid
func DeleteBandmember(bandmember structs.BandMember) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	_, err := db.Model(&bandmember).Where("bandmember.band_id = ?", bandmember.BandID).Where("bandmember.musician_id = ?", bandmember.MusicianID).Delete()
	return err
}
