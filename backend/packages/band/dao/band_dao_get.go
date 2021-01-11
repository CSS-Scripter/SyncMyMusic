package banddao

import (
	database "app/database"
	"app/structs"
)

// GetBandsUserIsIn queries for bands the user is in
func GetBandsUserIsIn(userID int) ([]structs.Band, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	result := []structs.Band{}

	subquery := db.Model(&structs.BandMember{}).Column("band_id").Where("musician_id = ?", userID)
	err := db.Model(&result).Where("id IN (?)", subquery).Select()

	return result, err
}

// GetBand queries for bands with the given ID
func GetBand(bandID int) (structs.Band, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	result := structs.Band{ID: bandID}

	err := db.Model(&result).WherePK().Select()
	return result, err
}

// GetJoinRequests get all join requests for a given band
func GetJoinRequests(bandID int) ([]structs.JoinRequest, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	result := []structs.JoinRequest{}
	err := db.Model(&result).Where("band_id = ?", bandID).Select()

	return result, err
}
