package songinstrumentdao

import (
	"app/database"
	"app/structs"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSongInstrumentDaoSuite struct {
	suite.Suite
	Owner          *structs.Musician
	Band           *structs.Band
	Song           *structs.Song
	Instrument     *structs.Instrument
	SongInstrument *structs.SongInstrument
}

func (suite *TestSongInstrumentDaoSuite) SetupTest() {
	suite.SongInstrument = &structs.SongInstrument{
		SongID:       999,
		InstrumentID: 999,
		PDFLocation:  "hier/staat/ie/dan.pdf",
	}

	suite.Instrument = &structs.Instrument{
		ID:     999,
		BandID: 999,
		Name:   "trompet",
	}

	suite.Song = &structs.Song{
		BandID: 999,
		ID:     999,
		Name:   "liedje",
		Artist: "artiest",
	}
	suite.Band = &structs.Band{
		ID:      999,
		Name:    "Adest musica",
		OwnerID: 999,
	}
	suite.Owner = &structs.Musician{
		MusicianID: 999,
		Name:       "Frits",
		Email:      "frits@email.nl",
		Password:   "geheim",
	}

	db := database.ConnectToDatabase()
	defer db.Close()

	db.Model(suite.Owner).Insert()
	db.Model(suite.Band).Insert()
	db.Model(suite.Song).Insert()
	db.Model(suite.Instrument).Insert()
}

func (suite *TestSongInstrumentDaoSuite) TearDownTest() {
	db := database.ConnectToDatabase()
	defer db.Close()

	_, err := db.Model(suite.SongInstrument).Where("song_id = ?", suite.SongInstrument.SongID).Delete()
	_, err = db.Model(suite.Instrument).WherePK().Delete()
	_, err = db.Model(suite.Song).WherePK().Delete()
	_, err = db.Model(suite.Band).WherePK().Delete()
	_, err = db.Model(suite.Owner).WherePK().Delete()
	fmt.Print(err)
}

func (suite *TestSongInstrumentDaoSuite) TestCreateSongInstrument_ShouldReturnListWithLengthOne() {
	err := CreateSongInstrument(suite.SongInstrument)

	suite.Nil(err, "There shouldn't be an error")

	db := database.ConnectToDatabase()
	defer db.Close()

	result := structs.SongInstrument{}

	err = db.Model(&result).Where("song_id = ?", suite.SongInstrument.SongID).Select()
	suite.Nil(err, "Something went wrong finding SongInstrument")
	suite.Equal(suite.SongInstrument.SongID, result.SongID)
}

func TestRunSongInstrumentDao(t *testing.T) {
	suite.Run(t, new(TestSongInstrumentDaoSuite))
}
