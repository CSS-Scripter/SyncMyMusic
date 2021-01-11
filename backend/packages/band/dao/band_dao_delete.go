package banddao

import (
	"app/database"
	"app/structs"

	"github.com/go-pg/pg/v10"
)

// DeleteBand checks if the band still has members left besides the owner himself. If so, it will give an error. Otherwise it will delete the band.
func DeleteBand(band *structs.Band) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	err := deleteBandRelations(db, band.ID)
	if err != nil {
		return err
	}

	_, err = db.Model(band).WherePK().Delete()
	return err
}

func deleteBandRelations(db *pg.DB, bandID int) error {
	deleteRelationSteps := []func(*pg.DB, int) error{
		deleteSongInstrumentsOfBand,
		deleteSongsOfBand,
		deleteMembersOfBand,
		deleteInstrumentsOfBand,
		deleteJoinRequestsOfBand,
	}

	for _, step := range deleteRelationSteps {
		err := step(db, bandID)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteSongInstrumentsOfBand(db *pg.DB, bandID int) error {
	subquery := db.Model(&structs.Instrument{}).Column("id").Where("band_id = ?", bandID)
	_, err := db.Model(&structs.SongInstrument{}).Where("instrument_id IN (?)", subquery).Delete()
	return err
}

func deleteSongsOfBand(db *pg.DB, bandID int) error {
	_, err := db.Model(&structs.Song{}).Where("band_id = ?", bandID).Delete()
	return err
}

func deleteMembersOfBand(db *pg.DB, bandID int) error {
	_, err := db.Model(&structs.BandMember{}).Where("band_id = ?", bandID).Delete()
	return err
}

func deleteInstrumentsOfBand(db *pg.DB, bandID int) error {
	_, err := db.Model(&structs.Instrument{}).Where("band_id = ?", bandID).Delete()
	return err
}

func deleteJoinRequestsOfBand(db *pg.DB, bandID int) error {
	_, err := db.Model(&structs.JoinRequest{BandID: bandID}).Where("band_id = ?", bandID).Delete()
	return err
}

// DeleteJoinRequest deletes a join request from the database
func DeleteJoinRequest(userID int, bandID int) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	_, err := db.Model(&structs.JoinRequest{}).Where("musician_id = ?", userID).Where("band_id = ?", bandID).Delete()
	return err
}
