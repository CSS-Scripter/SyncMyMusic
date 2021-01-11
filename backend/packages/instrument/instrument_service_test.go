package instrument

import (
	"app/database"
	"app/structs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestInstrumentServiceSuite struct {
	suite.Suite
	Instrument  *structs.Instrument
	Instrument2 *structs.Instrument
	Band        *structs.Band
	User        *structs.Musician
}

func (s *TestInstrumentServiceSuite) SetupTest() {
	s.Band = &structs.Band{
		ID:      666,
		Name:    "My Great Band",
		OwnerID: 800,
	}
	s.User = &structs.Musician{
		MusicianID: 800,
		Name:       "UserForInstrumentSrvc",
		Email:      "instrument@service.com",
		Password:   "test",
	}
	s.Instrument = &structs.Instrument{
		ID:     999,
		BandID: 666,
		Name:   "trompet",
	}
	s.Instrument2 = &structs.Instrument{
		ID:     888,
		BandID: 666,
		Name:   "trompet",
	}

	db := database.ConnectToDatabase()
	defer db.Close()

	db.Model(s.User).Insert()
	db.Model(s.Band).Insert()
	db.Model(s.Instrument).Insert()
}

func (s *TestInstrumentServiceSuite) TearDownTest() {
	db := database.ConnectToDatabase()
	defer db.Close()

	db.Model(s.Instrument).WherePK().Delete()
	db.Model(s.Instrument2).WherePK().Delete()

	db.Model(s.Band).WherePK().Delete()
	db.Model(s.User).WherePK().Delete()
}

func (s *TestInstrumentServiceSuite) Test_CreateInstrument() {
	err := CreateInstrument(s.Instrument2)
	s.Nil(err)
}

func (s *TestInstrumentServiceSuite) Test_GetInstrumentsForBand() {
	instruments, err := GetInstrumentsForBand(666)
	s.Nil(err)
	s.Greater(len(instruments), 0)
	s.Equal(instruments[0].Name, s.Instrument.Name)
}

func (s *TestInstrumentServiceSuite) Test_DeleteInstrument() {
	err := DeleteInstrument(999)
	s.Nil(err)
}

func TestRunInstrumentService(t *testing.T) {
	suite.Run(t, new(TestInstrumentServiceSuite))
}
