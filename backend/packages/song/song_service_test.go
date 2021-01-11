package song

import (
	banddao "app/packages/band/dao"
	instrumentdao "app/packages/instrument/dao"
	musiciandao "app/packages/musician/dao"
	songinstrumentdao "app/packages/songinstrument/dao"
	"app/structs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSongServiceSuite struct {
	suite.Suite
	User       *structs.Musician
	Band       *structs.Band
	Song       *structs.Song
	SongID     int
	Instrument *structs.Instrument
}

func (s *TestSongServiceSuite) SetupSuite() {
	user := &structs.Musician{
		MusicianID: 998,
		Name:       "Tester",
		Email:      "test@hello.com",
		Password:   "passwd",
	}
	musiciandao.CreateMusician(user)
	s.User = user

	band := &structs.Band{
		ID:      998,
		Name:    "Testers",
		OwnerID: user.MusicianID,
	}
	banddao.CreateBandForUser(user.MusicianID, band)
	s.Band = band

	instrument := &structs.Instrument{
		ID:     998,
		BandID: band.ID,
		Name:   "Guitar",
	}
	instrumentdao.CreateInstrument(instrument)
	s.Instrument = instrument
}

func (s *TestSongServiceSuite) TearDownSuite() {
	banddao.DeleteBand(s.Band)
	musiciandao.DeleteMusicianByID(s.User.MusicianID)
}

func (s *TestSongServiceSuite) SetupTest() {
	s.Song = &structs.Song{
		Name:   "TestSong",
		BandID: s.Band.ID,
		Artist: "Testers",
	}
}

func (s *TestSongServiceSuite) Test_A_CreateSongForBand() {
	songID, err := CreateSongForBand(s.Band.ID, s.Song)
	s.SongID = songID
	s.Nil(err)
	s.Greater(songID, 0)
}

func (s *TestSongServiceSuite) Test_B_GetSongsForBand() {
	songs, err := GetSongsForBand(s.Band.ID)
	s.Nil(err)
	s.Equal(len(songs), 1)
	s.Equal(songs[0].Name, s.Song.Name)
	s.Equal(songs[0].Artist, s.Song.Artist)
}

func (s *TestSongServiceSuite) Test_C_UpdateSong() {
	updatedSongTitle := s.Song.Name + "!"
	updatedSong := &structs.Song{
		ID:     s.SongID,
		Name:   updatedSongTitle,
		Artist: s.Song.Artist,
		BandID: s.Song.BandID,
	}
	err := UpdateSong(updatedSong)
	s.Nil(err)

	songs, _ := GetSongsForBand(s.Band.ID)
	s.Equal(updatedSongTitle, songs[0].Name)
}

func (s *TestSongServiceSuite) Test_D_GetInstrumentsForSong() {
	songInstrument := &structs.SongInstrument{
		SongID:       s.SongID,
		InstrumentID: s.Instrument.ID,
		PDFLocation:  "SomeFile.PDF",
	}
	songinstrumentdao.CreateSongInstrument(songInstrument)

	songInstruments, err := GetInstrumentsForSong(s.SongID)
	s.Nil(err)
	s.Equal(len(songInstruments), 1)
	s.Equal(songInstruments[0].PDFLocation, songInstrument.PDFLocation)
}

func (s *TestSongServiceSuite) Test_E_GetBandIDFromSong() {
	bandID, err := GetBandIDFromSong(s.SongID)
	s.Nil(err)
	s.Equal(bandID, s.Band.ID)
}

func (s *TestSongServiceSuite) Test_F_DeleteSong() {
	err := DeleteSong(s.SongID)
	s.Nil(err)
}

func TestRunSongServiceSuite(t *testing.T) {
	suite.Run(t, new(TestSongServiceSuite))
}
