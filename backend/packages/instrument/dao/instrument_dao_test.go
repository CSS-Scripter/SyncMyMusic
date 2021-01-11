package instrumentdao

import (
	"app/database"
	"app/structs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestInstrumentDaoSuite struct {
	suite.Suite
	Band        *structs.Band
	User        *structs.Musician
	Instrument1 *structs.Instrument
	Instrument2 *structs.Instrument
	Instrument3 *structs.Instrument
}

func (suite *TestInstrumentDaoSuite) SetupTest() {
	suite.Band = &structs.Band{
		ID:      990,
		Name:    "tromphia",
		OwnerID: 900,
	}
	suite.User = &structs.Musician{
		MusicianID: 900,
		Name:       "UserForInstrumentDao",
		Email:      "instrument@dao.com",
		Password:   "test",
	}
	suite.Instrument1 = &structs.Instrument{
		ID:     998,
		BandID: 990,
		Name:   "trompet",
	}
	suite.Instrument2 = &structs.Instrument{
		ID:     999,
		BandID: 990,
		Name:   "trombone",
	}
	suite.Instrument3 = &structs.Instrument{
		ID:     111,
		BandID: 990,
		Name:   "trommel",
	}

	//musician, _ = musiciandao.CreateMusician(musician)

	db := database.ConnectToDatabase()
	defer db.Close()

	db.Model(suite.User).Insert()
	db.Model(suite.Band).Insert()

	db.Model(suite.Instrument1).Insert()
	db.Model(suite.Instrument2).Insert()
}

func (suite *TestInstrumentDaoSuite) TearDownTest() {
	db := database.ConnectToDatabase()
	defer db.Close()

	db.Model(suite.Instrument1).WherePK().Delete()
	db.Model(suite.Instrument2).WherePK().Delete()

	db.Model(suite.Band).WherePK().Delete()
	db.Model(suite.User).WherePK().Delete()
}

func (suite *TestInstrumentDaoSuite) TestGetInstruments_ShouldReturnListWithLengthTwo() {
	bandID := 990
	expected := 2

	instruments, err := GetInstrumentsForBand(bandID)

	suite.Nil(err, "There shouldn't be an error")
	suite.Equal(expected, len(instruments), "instruments should be of length 2")
}

func (suite *TestInstrumentDaoSuite) TestCreateInstrument_GetInstrumentsShouldReturnListWithLengthThree() {
	BandID := 990
	expectedLength := 3

	err := CreateInstrument(suite.Instrument3)

	suite.Nil(err, "There shouldn't be an error")
	instruments, _ := GetInstrumentsForBand(BandID)
	suite.Equal(expectedLength, len(instruments))
}

func (suite *TestInstrumentDaoSuite) TestDeleteInstrument_GetInstrumentsShouldReturnListWithLengthTwo() {
	BandID := 990
	instrumentToDeleteID := 111
	expectedLength := 2

	err := DeleteInstrument(instrumentToDeleteID)

	suite.Nil(err, "There shouldn't be an error")
	instruments, _ := GetInstrumentsForBand(BandID)
	suite.Equal(expectedLength, len(instruments))
}

func TestRunInstrumentDao(t *testing.T) {
	suite.Run(t, new(TestInstrumentDaoSuite))
}
