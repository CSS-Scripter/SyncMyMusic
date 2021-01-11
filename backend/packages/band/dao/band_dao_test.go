package banddao

import (
	musiciandao "app/packages/musician/dao"
	"app/structs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestBandDaoSuite struct {
	suite.Suite
	User   *structs.Musician
	Band   *structs.Band
	BandID int
}

func (suite *TestBandDaoSuite) SetupSuite() {
	musician := &structs.Musician{
		MusicianID: 997,
		Name:       "Tester",
		Email:      "test@hello.com",
		Password:   "passwd",
	}
	musician, _ = musiciandao.CreateMusician(musician)
	suite.User = musician
}

func (suite *TestBandDaoSuite) TearDownSuite() {
	musiciandao.DeleteMusicianByID(suite.User.MusicianID)
}

func (suite *TestBandDaoSuite) SetupTest() {
	suite.Band = &structs.Band{
		Name:    "Hello World",
		OwnerID: suite.User.MusicianID,
	}
}

func (suite *TestBandDaoSuite) TestCreateBand() {
	bands, errForGet := GetBandsUserIsIn(suite.User.MusicianID)
	amountOfBands := len(bands)
	bandID, errForCreate := CreateBandForUser(suite.User.MusicianID, suite.Band)
	bands, errForGetTwo := GetBandsUserIsIn(suite.User.MusicianID)

	suite.BandID = bandID
	createdBand := suite.Band
	createdBand.ID = bandID

	suite.Nil(errForCreate, "There shouldn't be an error")
	suite.Nil(errForGet, "There shouldn't be an error")
	suite.Nil(errForGetTwo, "There shouldn't be an error")
	suite.Equal(len(bands), amountOfBands+1, "There should be 1 band")
	suite.Equal(createdBand.Name, suite.Band.Name, "The name should be the same as given name")
	suite.Equal(createdBand.OwnerID, suite.User.MusicianID, "The given UserID should be the owner")
}

func (suite *TestBandDaoSuite) TestDeleteBand() {
	bands, errForGet := GetBandsUserIsIn(suite.User.MusicianID)
	amountOfBands := len(bands)
	suite.Band.ID = suite.BandID
	errForDelete := DeleteBand(suite.Band)
	bands, errForGetTwo := GetBandsUserIsIn(suite.User.MusicianID)

	suite.Nil(errForGet, "There shouldn't be an error")
	suite.Nil(errForDelete, "There shouldn't be an error")
	suite.Nil(errForGetTwo, "There shouldn't be an error")
	suite.Equal(len(bands), amountOfBands-1, "There should be 1 less band")
}

func TestRunBandDao(t *testing.T) {
	suite.Run(t, new(TestBandDaoSuite))
}
