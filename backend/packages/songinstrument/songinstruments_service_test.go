package songinstrument

import (
	"app/structs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSongInstrumentServiceSuite struct {
	suite.Suite
}

func (suite *TestSongInstrumentServiceSuite) TestCreateSongInstrument_ShouldCreateInstrumentAsExpected() {
	SongID := 1
	InstrumentID := 1
	BandID := 1
	PDFLocation := "./files/1/1/1.pdf"
	expected := &structs.SongInstrument{
		SongID:       SongID,
		InstrumentID: InstrumentID,
		PDFLocation:  PDFLocation,
	}

	oldSongGetBandIDFromSong := songGetBandIDFromSong
	oldDaoCreateSongInstrument := daoCreateSongInstrument
	defer func() {
		songGetBandIDFromSong = oldSongGetBandIDFromSong
		daoCreateSongInstrument = oldDaoCreateSongInstrument
	}()

	songGetBandIDFromSong = func(int) (int, error) {
		return BandID, nil
	}
	result := &structs.SongInstrument{}
	daoCreateSongInstrument = func(songInstrument *structs.SongInstrument) error {
		result = songInstrument
		return nil
	}

	returnedBandID, err := CreateSongInstrument(SongID, InstrumentID)

	suite.Nil(err, "There shouldn't be an error")
	suite.Equal(BandID, returnedBandID, "Should return bandID 1")
	suite.Equal(expected, result, "daoCreateSongInstrument Not called with expected instrument")
}

func TestRunSongInstrumentServiceSuite(t *testing.T) {
	suite.Run(t, new(TestSongInstrumentServiceSuite))
}
