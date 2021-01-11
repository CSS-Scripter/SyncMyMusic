package bandmemberdao

import (
	"app/database"
	"app/structs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestBandmemberDaoSuite struct {
	suite.Suite
	Band        *structs.Band
	User1       *structs.Musician
	User2       *structs.Musician
	BandMember1 *structs.BandMember
	BandMember2 *structs.BandMember
	Instrument1 *structs.Instrument
	Instrument2 *structs.Instrument
}

func (suite *TestBandmemberDaoSuite) SetupTest() {
	suite.User1 = &structs.Musician{
		MusicianID: 998,
		Name:       "Richard",
		Email:      "hallo@email.nl",
		Password:   "geheim",
	}
	suite.User2 = &structs.Musician{
		MusicianID: 999,
		Name:       "Frits",
		Email:      "frits@email.nl",
		Password:   "geheim",
	}
	suite.Band = &structs.Band{
		ID:      999,
		Name:    "Adest musica",
		OwnerID: 999,
	}
	suite.BandMember1 = &structs.BandMember{
		MusicianID: 998,
		BandID:     999,
	}
	suite.BandMember2 = &structs.BandMember{
		MusicianID: 999,
		BandID:     999,
	}
	suite.Instrument1 = &structs.Instrument{
		ID:     998,
		BandID: 999,
		Name:   "trompet",
	}
	suite.Instrument2 = &structs.Instrument{
		ID:     999,
		BandID: 999,
		Name:   "trombone",
	}

	db := database.ConnectToDatabase()
	defer db.Close()

	db.Model(suite.User1).Insert()
	db.Model(suite.User2).Insert()

	db.Model(suite.Band).Insert()

	db.Model(suite.BandMember1).Insert()
	db.Model(suite.BandMember2).Insert()

	db.Model(suite.Instrument1).Insert()
	db.Model(suite.Instrument2).Insert()
}

func (suite *TestBandmemberDaoSuite) TearDownTest() {
	db := database.ConnectToDatabase()
	defer db.Close()

	db.Model(suite.BandMember1).Where("band_id = ?", suite.BandMember1.BandID).Delete()
	db.Model(suite.BandMember2).Where("band_id = ?", suite.BandMember2.BandID).Delete()

	db.Model(suite.Instrument1).WherePK().Delete()
	db.Model(suite.Instrument2).WherePK().Delete()

	db.Model(suite.Band).WherePK().Delete()

	db.Model(suite.User1).WherePK().Delete()
	db.Model(suite.User2).WherePK().Delete()
}

func (suite *TestBandmemberDaoSuite) TestGetBandmembers_ShouldReturnListWithLengthTwo() {
	bandID := 999
	expected := 2

	members, err := GetBandmembers(bandID)

	suite.Nil(err, "There shouldn't be an error")
	suite.Equal(expected, len(members), "members should be of length 2")
}

func (suite *TestBandmemberDaoSuite) TestUpdateBandmember_ShouldReturnMemberWithInstrumentOne() {
	expectedInstrument := suite.Instrument1
	MusicianID := 998
	BandID := 999
	InstrumentID := 998
	updatedMember := &structs.BandMember{
		MusicianID:   MusicianID,
		BandID:       BandID,
		InstrumentID: InstrumentID,
	}

	err := UpdateBandmember(*updatedMember)

	suite.Nil(err, "There shouldn't be an error")
	members, _ := GetBandmembers(BandID)
	for _, member := range members {
		if member.Musician.MusicianID == MusicianID {
			suite.Equal(expectedInstrument, member.Instrument, "Instrument not added")
		}
	}
}

func (suite *TestBandmemberDaoSuite) TestDeleteBandmember_GetBandmembersShouldReturnListWithLengthOne() {
	BandID := 999
	bandmemberToDelete := suite.BandMember2
	expectedLength := 1

	err := DeleteBandmember(*bandmemberToDelete)

	suite.Nil(err, "There shouldn't be an error")
	members, _ := GetBandmembers(BandID)
	suite.Equal(expectedLength, len(members))
}

func TestRunBandDao(t *testing.T) {
	suite.Run(t, new(TestBandmemberDaoSuite))
}
