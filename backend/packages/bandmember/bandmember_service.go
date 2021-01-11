package bandmember

import (
	dao "app/packages/bandmember/dao"
	"app/structs"
)

// GetBandmembers return a list of bandmembers
func GetBandmembers(bandID int) ([]structs.BandMember, error) {
	return dao.GetBandmembers(bandID)
}

// UpdateBandmember updates a single bandmember
func UpdateBandmember(member structs.BandMember) error {
	return dao.UpdateBandmember(member)
}

// DeleteBandmember delets a single bandmember
func DeleteBandmember(member structs.BandMember) error {
	return dao.DeleteBandmember(member)
}
