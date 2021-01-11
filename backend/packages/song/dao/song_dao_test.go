package songdao

import (
	"app/packages/band"
	musiciandao "app/packages/musician/dao"
	"app/structs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSongDaoSuite struct {
	suite.Suite
	BandID     int
	Song       *structs.Song
	Band       *structs.Band
	Instrument *structs.Instrument
	SongID     int
	User       *structs.Musician
}

func (suite *TestSongDaoSuite) SetupSuite() {
	user := &structs.Musician{
		MusicianID: 997,
		Name:       "Tester",
		Email:      "test@hello.com",
		Password:   "passwd",
	}
	user, _ = musiciandao.CreateMusician(user)
	suite.User = user

	testBand := &structs.Band{
		ID:   998,
		Name: "Testers",
	}
	band.CreateBandForUser(user.MusicianID, testBand)
	suite.BandID = testBand.ID
}

func (suite *TestSongDaoSuite) TearDownSuite() {
	band.DeleteBand(suite.BandID, suite.User.Email)
	musiciandao.DeleteMusicianByID(suite.User.MusicianID)
}

func (suite *TestSongDaoSuite) SetupTest() {
	suite.Song = &structs.Song{
		Name:   "Clouds",
		Artist: "Paper Idol",
		BandID: suite.BandID,
	}
	suite.Band = &structs.Band{
		ID:      999,
		Name:    "Adest musica",
		OwnerID: 999,
	}
	suite.Instrument = &structs.Instrument{
		ID:     999,
		BandID: 999,
		Name:   "trompet",
	}
}

func (suite *TestSongDaoSuite) Test_A_CreateSong() {
	ID, errForCreate := CreateSongForBand(suite.Song)
	suite.SongID = ID
	song, errForGet := GetSong(ID)

	suite.Nil(errForCreate, "There shouldn't be an error")
	suite.Nil(errForGet, "There shouldn't be an error")
	suite.Equal(song.Name, suite.Song.Name, "This song should have been added")
}

func (suite *TestSongDaoSuite) Test_B_UpdateSong() {
	song, _ := GetSong(suite.SongID)
	songName := song.Name
	song.Name += "!"
	errForUpdate := UpdateSong(&song)
	updatedSong, _ := GetSong(suite.SongID)

	suite.Nil(errForUpdate)
	suite.Equal(songName+"!", updatedSong.Name)
}

func (suite *TestSongDaoSuite) Test_C_DeleteSong() {
	songs, errForGet := GetSongsForBand(suite.BandID)
	amountOfSongs := len(songs)
	errForDelete := DeleteSong(suite.SongID)
	songs, errForGetTwo := GetSongsForBand(suite.BandID)

	suite.Nil(errForGet)
	suite.Nil(errForDelete)
	suite.Nil(errForGetTwo)

	suite.Equal(amountOfSongs-1, len(songs))
}

func (suite *TestSongDaoSuite) TestUpdateSong() {
	ID, _ := CreateSongForBand(suite.Song)
	song, _ := GetSong(ID)
	songName := song.Name
	song.Name += "!"
	errForUpdate := UpdateSong(&song)
	updatedSong, _ := GetSong(ID)

	suite.Nil(errForUpdate)
	suite.Equal(songName+"!", updatedSong.Name)
}

func TestRunSongDao(t *testing.T) {
	suite.Run(t, new(TestSongDaoSuite))
}
