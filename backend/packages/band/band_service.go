package band

import (
	"app/errors"
	dao "app/packages/band/dao"
	"app/packages/musician"
	"app/structs"
	"fmt"
	"os"
)

// GetBand requests a single band with the given ID
func GetBand(bandID int) (structs.Band, error) {
	return dao.GetBand(bandID)
}

// AddUserToBand adds an existing user to band by email
func AddUserToBand(bandID int, musicianObj *structs.Musician) error {
	email := musicianObj.Email
	var retrievedMusician, _ = musician.GetMusicianByEmail(email)
	bandmember := &structs.BandMember{MusicianID: retrievedMusician.MusicianID, BandID: bandID}
	return dao.AddUserToBand(bandmember)
}

// GetBandsUserIsIn requests the bands for a certain user
func GetBandsUserIsIn(userID int) ([]structs.Band, error) {
	return dao.GetBandsUserIsIn(userID)
}

// CreateBandForUser creates an band for the user
func CreateBandForUser(userID int, band *structs.Band) (int, error) {
	band.OwnerID = userID
	return dao.CreateBandForUser(userID, band)
}

// DeleteBand deletes an band
func DeleteBand(bandID int, email string) (int, error) {
	isOwner, err := checkIfUserIsBandOwner(email, bandID)
	if err != nil {
		return 500, err
	}
	if isOwner {
		band := &structs.Band{ID: bandID}
		err = dao.DeleteBand(band)
		if err != nil {
			return 500, err
		}
		return 200, os.RemoveAll(fmt.Sprintf("files/%v", bandID))
	}
	return 401, nil
}

// JoinRequest registers a join request for a certain band
func JoinRequest(email string, bandID int) error {
	user, err := musician.GetMusicianByEmail(email)
	if err != nil {
		return err
	}
	_, err = GetBand(bandID)
	if err != nil {
		return new(errors.BandDoesNotExistError)
	}
	return dao.CreateJoinRequest(user.MusicianID, bandID)
}

func checkIfUserIsBandOwner(email string, bandID int) (bool, error) {
	user, err := musician.GetMusicianByEmail(email)
	if err != nil {
		return false, err
	}
	band, err := GetBand(bandID)
	if err != nil {
		return false, err
	}
	return band.OwnerID == user.MusicianID, nil
}

// GetJoinRequestsForBand returns all found join requests for a band
func GetJoinRequestsForBand(email string, bandID int) (map[int]string, int, error) {
	isOwner, err := checkIfUserIsBandOwner(email, bandID)
	if err != nil {
		return nil, 500, err
	}
	if isOwner {
		result, err := dao.GetJoinRequests(bandID)
		returnMap := map[int]string{}
		for _, res := range result {
			user, err := musician.GetMusicianByID(res.UserID)
			if err != nil {
				return nil, 500, err
			}
			returnMap[res.UserID] = user.Name
		}
		return returnMap, 200, err
	}
	return map[int]string{}, 401, nil
}

// AcceptJoinRequest accepts a join request.
// Returns a statuscode and error when one occured
func AcceptJoinRequest(email string, bandID int, userID int) (int, error) {
	isOwner, err := checkIfUserIsBandOwner(email, bandID)
	if err != nil || !isOwner {
		return 401, err
	}
	err = dao.DeleteJoinRequest(userID, bandID)
	if err != nil {
		return 500, err
	}
	err = dao.AddUserToBand(&structs.BandMember{MusicianID: userID, BandID: bandID})
	if err != nil {
		return 500, err
	}
	return 200, nil
}

// DenyJoinRequest accepts a join request.
// Returns a statuscode and error when one occured
func DenyJoinRequest(email string, bandID int, userID int) (int, error) {
	isOwner, err := checkIfUserIsBandOwner(email, bandID)
	if err != nil || !isOwner {
		return 401, err
	}
	err = dao.DeleteJoinRequest(userID, bandID)
	if err != nil {
		return 500, err
	}
	return 200, nil
}
