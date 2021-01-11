package musiciandao

import (
	"app/structs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestMusicianDaoSuite struct {
	suite.Suite
	Musician   *structs.Musician
	MusicianID int
}

func (s *TestMusicianDaoSuite) SetupTest() {
	s.Musician = &structs.Musician{
		Email:    "Hi@myhost.app",
		Password: "Hi",
		Name:     "Thom",
	}
}

func (s *TestMusicianDaoSuite) Test_A_CreateMusician() {
	result, creationErr := CreateMusician(s.Musician)
	s.MusicianID = result.MusicianID
	s.Nil(creationErr)
	s.Greater(result.MusicianID, 0)
}

func (s *TestMusicianDaoSuite) Test_B_GetMusicianByEmail() {
	result, err := GetMusicianByEmail(s.Musician.Email)
	s.Nil(err)
	s.Equal(result.Name, s.Musician.Name)
	s.Equal(result.Password, s.Musician.Password)
	s.Equal(result.MusicianID, s.MusicianID)
}

func (s *TestMusicianDaoSuite) Test_C_GetMusicianByID() {
	result, err := GetMusicianByID(s.MusicianID)
	s.Nil(err)
	s.Equal(result.Name, s.Musician.Name)
	s.Equal(result.Email, s.Musician.Email)
}

func (s *TestMusicianDaoSuite) Test_D_DeleteMusicianByID() {
	err := DeleteMusicianByID(s.MusicianID)
	s.Nil(err)
}

func TestRunMusicianDao(t *testing.T) {
	suite.Run(t, new(TestMusicianDaoSuite))
}
