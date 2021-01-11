package band

import (
	"app/packages/bandmember"
	musiciandao "app/packages/musician/dao"
	"app/structs"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestBandServiceSuite struct {
	suite.Suite
	Band       *structs.Band
	User       *structs.Musician
	SecondUser *structs.Musician
	BandID     int
}

func (s *TestBandServiceSuite) SetupSuite() {
	firstUser := &structs.Musician{
		Name:       "Frits",
		Email:      "frits@email.nl",
		MusicianID: 996,
		Password:   "geheim",
	}
	musiciandao.CreateMusician(firstUser)
	s.User = firstUser

	secondUser := &structs.Musician{
		Name:       "Tester",
		Email:      "test@hello.com",
		MusicianID: 995,
		Password:   "passwd",
	}
	musiciandao.CreateMusician(secondUser)
	s.SecondUser = secondUser
}

func (s *TestBandServiceSuite) TearDownSuite() {
	musiciandao.DeleteMusicianByID(s.User.MusicianID)
	musiciandao.DeleteMusicianByID(s.SecondUser.MusicianID)
}

func (s *TestBandServiceSuite) SetupTest() {
	s.Band = &structs.Band{
		Name: "Hello World",
	}
}

func (s *TestBandServiceSuite) Test_A_CreateBandForUser() {
	bandID, err := CreateBandForUser(s.User.MusicianID, s.Band)
	s.BandID = bandID
	s.Nil(err)
	s.Greater(bandID, 0)
}

func (s *TestBandServiceSuite) Test_B_BandOwnership() {
	isOwner, err := checkIfUserIsBandOwner(s.User.Email, s.BandID)
	s.Nil(err)
	s.True(isOwner)
}

func (s *TestBandServiceSuite) Test_C_GetBand() {
	band, err := GetBand(s.BandID)
	s.Nil(err)
	s.Equal(band.Name, s.Band.Name)
}

func (s *TestBandServiceSuite) Test_D_AddUserToBand() {
	err := AddUserToBand(s.BandID, s.SecondUser)
	s.Nil(err)
}

func (s *TestBandServiceSuite) Test_E_GetBandsUserIsIn() {
	bands, err := GetBandsUserIsIn(s.SecondUser.MusicianID)
	s.Nil(err)
	s.Greater(len(bands), 0)
	s.Equal(bands[0].Name, s.Band.Name)
}

func (s *TestBandServiceSuite) Test_F_RemoveBandMember() {
	err := bandmember.DeleteBandmember(structs.BandMember{BandID: s.BandID, MusicianID: s.SecondUser.MusicianID})
	bands, _ := GetBandsUserIsIn(s.SecondUser.MusicianID)
	s.Nil(err)
	s.Equal(bands, []structs.Band{})
}

func (s *TestBandServiceSuite) Test_G_CreateJoinRequest() {
	err := JoinRequest(s.SecondUser.Email, s.BandID)
	s.Nil(err)
}

func (s *TestBandServiceSuite) Test_H_GetJoinRequests() {
	requests, status, err := GetJoinRequestsForBand(s.User.Email, s.BandID)
	s.Nil(err)
	s.Equal(len(requests), 1)
	s.Equal(requests[s.SecondUser.MusicianID], s.SecondUser.Name)
	s.Equal(status, 200)
}

func (s *TestBandServiceSuite) Test_I_UnauthorizedGetJoinRequests() {
	requests, status, err := GetJoinRequestsForBand(s.SecondUser.Email, s.BandID)
	s.Nil(err)
	s.Equal(status, 401)
	s.Equal(requests, map[int]string{})
}

func (s *TestBandServiceSuite) Test_J_UnauthorizedDenyJoinRequest() {
	status, err := DenyJoinRequest(s.SecondUser.Email, s.BandID, s.SecondUser.MusicianID)
	s.Nil(err)
	s.Equal(status, 401)
}

func (s *TestBandServiceSuite) Test_K_DenyJoinRequests() {
	status, err := DenyJoinRequest(s.User.Email, s.BandID, s.SecondUser.MusicianID)
	s.Nil(err)
	s.Equal(status, 200)
}

func (s *TestBandServiceSuite) Test_L_JoinRequestDenied() {
	requests, status, err := GetJoinRequestsForBand(s.User.Email, s.BandID)
	s.Nil(err)
	s.Equal(len(requests), 0)
	s.Equal(status, 200)

	bands, err := GetBandsUserIsIn(s.SecondUser.MusicianID)
	s.Nil(err)
	s.Equal(len(bands), 0)
}

func (s *TestBandServiceSuite) Test_M_SecondJoinRequest() {
	err := JoinRequest(s.SecondUser.Email, s.BandID)
	s.Nil(err)
}

func (s *TestBandServiceSuite) Test_N_UnauthorizedAcceptJoinRequest() {
	status, err := AcceptJoinRequest(s.SecondUser.Email, s.BandID, s.SecondUser.MusicianID)
	s.Nil(err)
	s.Equal(status, 401)
}

func (s *TestBandServiceSuite) Test_O_AcceptJoinRequest() {
	status, err := AcceptJoinRequest(s.User.Email, s.BandID, s.SecondUser.MusicianID)
	s.Nil(err)
	s.Equal(status, 200)
}

func (s *TestBandServiceSuite) Test_P_UnauthorizedDeleteBand() {
	status, err := DeleteBand(s.BandID, s.SecondUser.Email)
	s.Nil(err)
	s.Equal(status, 401)
}

func (s *TestBandServiceSuite) Test_Q_DeleteBand() {
	status, err := DeleteBand(s.BandID, s.User.Email)
	s.Nil(err)
	s.Equal(status, 200)
}

func TestRunBandService(t *testing.T) {
	suite.Run(t, new(TestBandServiceSuite))
}
