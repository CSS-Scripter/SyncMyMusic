package banddao

import (
	"app/database"
	"app/structs"
)

// CreateBandForUser creates a band with the given userID as owner
func CreateBandForUser(userID int, band *structs.Band) (int, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	_, err := db.Model(band).Insert()
	if err != nil {
		return 0, err
	}

	bandMember := &structs.BandMember{
		BandID:     band.ID,
		MusicianID: userID,
	}

	_, err = db.Model(bandMember).Insert()
	return band.ID, err
}

// AddUserToBand adds a new bandmember to a band
func AddUserToBand(bandmember *structs.BandMember) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	bandMember := &structs.BandMember{
		BandID:     bandmember.BandID,
		MusicianID: bandmember.MusicianID,
	}
	var err error
	_, err = db.Model(bandMember).Insert()
	return err
}

// CreateJoinRequest creates a join request
func CreateJoinRequest(userID int, bandID int) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	joinRequest := &structs.JoinRequest{
		UserID: userID,
		BandID: bandID,
	}
	var err error
	_, err = db.Model(joinRequest).Insert()
	return err
}
