package main

import (
	"app/packages/band"
	banddao "app/packages/band/dao"
	"app/packages/musician"
	musiciandao "app/packages/musician/dao"
	"app/packages/song"
	songdao "app/packages/song/dao"
	"app/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
)

type TestAppSuite struct {
	suite.Suite
	ATGetRoutes    []string
	ATUserName     string
	ATUserEmail    string
	ATUserPassword string
	ATUserID       int
	ATBandName     string
	ATBandID       int
	ATSongName     string
	ATSongArtist   string
	ATSongID       int
	ATTimeout      int
}

func (s *TestAppSuite) TearDownSuite() {
	songdao.DeleteSong(s.ATSongID)
	banddao.DeleteBand(&structs.Band{ID: s.ATBandID})
	musiciandao.DeleteMusicianByID(s.ATUserID)
}

func (s *TestAppSuite) SetupTest() {
	s.ATGetRoutes = []string{
		"/users/999/rehearsals",
		"/users/999/bands",
		"/bands/999",
		"/bands/999/request",
		"/bands/999/songs",
		"/songs/999",
		"/bands/999/instruments",
		"/songs/999/instruments",
		"/files/999/999/999",
		"/bands/999/members",
		"/members/hi@hi.nl",
	}
	s.ATUserName = "Tester"
	s.ATUserEmail = "test@hello.com"
	s.ATUserPassword = "passwd"
	s.ATBandName = "Testers"
	s.ATSongName = "Clouds"
	s.ATSongArtist = "PaperIdols"
	s.ATTimeout = 5000
}

func (s *TestAppSuite) Test_A_AppTest() {
	app := fiber.New()
	setupMiddleware(app)
	setupRoutes(app)

	for i := 0; i < len(s.ATGetRoutes); i++ {
		req := httptest.NewRequest("GET", s.ATGetRoutes[i], nil)
		resp, _ := app.Test(req)

		s.Equal(401, resp.StatusCode)
	}

}

func (s *TestAppSuite) Test_A_RouteRegister() {
	app := fiber.New()
	app.Post("/register", musician.RouteCreateMusician)

	requestJSON := "{" +
		"\"username\":\"" + s.ATUserName + "\"," +
		"\"email\":\"" + s.ATUserEmail + "\"," +
		"\"Password\":\"" + s.ATUserPassword + "\"" +
		"}"
	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(requestJSON))
	req.Header.Set("Content-Length", strconv.Itoa(len(requestJSON)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(resp.StatusCode, 200)
}

func (s *TestAppSuite) Test_A_RouteRegister_Error() {
	app := fiber.New()
	app.Post("/register", musician.RouteCreateMusician)

	requestJSON := "{" +
		"\"usme\":\"" + s.ATUserName + "\"," +
		"\"email\":\"" + s.ATUserEmail + "\"," +
		"\"Password\":\"" + s.ATUserPassword + "\"" +
		"}"
	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(requestJSON))
	req.Header.Set("Content-Length", strconv.Itoa(len(requestJSON)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(resp.StatusCode, 500)
}

func (s *TestAppSuite) Test_B_RouteGetMusician() {
	app := fiber.New()
	setupAuth(app)
	app.Get("/login/:email", musician.RouteGetMusician)

	req, _ := http.NewRequest(http.MethodGet, "/login/"+s.ATUserEmail, nil)
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)

	resp, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(resp.StatusCode, 200)

	body, _ := ioutil.ReadAll(resp.Body)
	bodyObject := parseResponseBodyToObject(body)

	s.ATUserID, err = strconv.Atoi(fmt.Sprintf("%v", bodyObject["MusicianID"]))

	s.Equal(s.ATUserEmail, fmt.Sprintf("%v", bodyObject["email"]))
	s.Equal(s.ATUserName, fmt.Sprintf("%v", bodyObject["username"]))
}

func (s *TestAppSuite) Test_B_RouteGetMusician_Error() {
	app := fiber.New()
	setupAuth(app)
	app.Get("/login/:email", musician.RouteGetMusician)

	req, _ := http.NewRequest(http.MethodGet, "/login/fake@email.com", nil)
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)

	resp, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(resp.StatusCode, 500)
}

func (s *TestAppSuite) Test_C_RouteCreateBand() {
	app := fiber.New()
	setupAuth(app)
	app.Post("/users/:userid/bands", band.RouteCreateBand)
	path := fmt.Sprintf("/users/%v/bands", s.ATUserID)

	requestJSON := `{"name":"` + s.ATBandName + `"}`
	req, _ := http.NewRequest(http.MethodPost, path, strings.NewReader(requestJSON))
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)
	req.Header.Set("Content-Length", strconv.Itoa(len(requestJSON)))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(res.StatusCode, 200)

	body, _ := ioutil.ReadAll(res.Body)
	bandID, err := strconv.Atoi(string(body))
	s.ATBandID = bandID

	s.Greater(bandID, 0)
}

func (s *TestAppSuite) Test_C_RouteCreateBand_Error() {
	app := fiber.New()
	setupAuth(app)
	app.Post("/users/:userid/bands", band.RouteCreateBand)
	path := fmt.Sprintf("/users/%v/bands", "Hello")

	requestJSON := `{"name":"` + s.ATBandName + `"}`
	req, _ := http.NewRequest(http.MethodPost, path, strings.NewReader(requestJSON))
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)
	req.Header.Set("Content-Length", strconv.Itoa(len(requestJSON)))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(res.StatusCode, 500)
}

func (s *TestAppSuite) Test_D_RouteGetSingleBand() {
	app := fiber.New()
	setupAuth(app)
	app.Get("/bands/:bandid", band.RouteGetSingleBand)
	path := fmt.Sprintf("/bands/%v", s.ATBandID)

	req, _ := http.NewRequest(http.MethodGet, path, nil)
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)

	res, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(res.StatusCode, 200)

	body, _ := ioutil.ReadAll(res.Body)
	bodyObject := parseResponseBodyToObject(body)

	s.Equal(bodyObject["name"], s.ATBandName)
	s.EqualValues(bodyObject["ID"], s.ATBandID)
	s.EqualValues(bodyObject["OwnerID"], s.ATUserID)
}

func (s *TestAppSuite) Test_D_RouteGetSingleBand_Error() {
	app := fiber.New()
	setupAuth(app)
	app.Get("/bands/:bandid", band.RouteGetSingleBand)
	path := fmt.Sprintf("/bands/%v", "Hello")

	req, _ := http.NewRequest(http.MethodGet, path, nil)
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)

	res, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(res.StatusCode, 500)
}

func (s *TestAppSuite) Test_E_RouteGetBandsUserIsIn() {
	app := fiber.New()
	setupAuth(app)
	app.Get("/users/:userid/bands/", band.RouteGetBandsUserIsIn)
	path := fmt.Sprintf("/users/%v/bands/", s.ATUserID)

	req, _ := http.NewRequest(http.MethodGet, path, nil)
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)

	res, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(res.StatusCode, 200)

	body, _ := ioutil.ReadAll(res.Body)
	bodyObject := parseResponseBodyToArray(body)
	band := bodyObject[0]

	s.Equal(band["name"], s.ATBandName)
	s.EqualValues(band["ID"], s.ATBandID)
	s.EqualValues(band["OwnerID"], s.ATUserID)
}

func (s *TestAppSuite) Test_E_RouteGetBandsUserIsIn_Error() {
	app := fiber.New()
	setupAuth(app)
	app.Get("/users/:userid/bands/", band.RouteGetBandsUserIsIn)
	path := fmt.Sprintf("/users/%v/bands/", "NotAnID")

	req, _ := http.NewRequest(http.MethodGet, path, nil)
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)

	res, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(res.StatusCode, 500)
}

func (s *TestAppSuite) Test_F_RouteCreateSong() {
	app := fiber.New()
	setupAuth(app)
	app.Post("/bands/:bandid/songs/", song.RouteCreateSongForBand)
	path := fmt.Sprintf("/bands/%v/songs/", s.ATBandID)

	requestJSON := `{"name":"` + s.ATSongName + `", "artist":"` + s.ATSongArtist + `"}`
	req, _ := http.NewRequest(http.MethodPost, path, strings.NewReader(requestJSON))
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)
	req.Header.Set("Content-Length", strconv.Itoa(len(requestJSON)))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(res.StatusCode, 200)

	body, _ := ioutil.ReadAll(res.Body)
	ID, err := strconv.Atoi(string(body))
	s.ATSongID = ID

	s.Greater(ID, 0)
}

func (s *TestAppSuite) Test_F_RouteCreateSong_Error() {
	app := fiber.New()
	setupAuth(app)
	app.Post("/bands/:bandid/songs/", song.RouteCreateSongForBand)
	path := fmt.Sprintf("/bands/%v/songs/", "NotAnID")

	requestJSON := `{"name":"` + s.ATSongName + `", "artist":"` + s.ATSongArtist + `"}`
	req, _ := http.NewRequest(http.MethodPost, path, strings.NewReader(requestJSON))
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)
	req.Header.Set("Content-Length", strconv.Itoa(len(requestJSON)))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(res.StatusCode, 500)
}

func (s *TestAppSuite) Test_G_RouteGetSongsForBand() {
	app := fiber.New()
	setupAuth(app)
	app.Get("/bands/:bandid/songs/", song.RouteGetSongsForBand)
	path := fmt.Sprintf("/bands/%v/songs/", s.ATBandID)

	req, _ := http.NewRequest(http.MethodGet, path, nil)
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)

	res, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(res.StatusCode, 200)

	body, _ := ioutil.ReadAll(res.Body)
	bodyObject := parseResponseBodyToArray(body)
	song := bodyObject[0]

	s.Equal(song["name"], s.ATSongName)
	s.Equal(song["artist"], s.ATSongArtist)
	s.EqualValues(song["ID"], s.ATSongID)
	s.EqualValues(song["BandID"], s.ATBandID)
}

func (s *TestAppSuite) Test_G_RouteGetSongsForBand_Error() {
	app := fiber.New()
	setupAuth(app)
	app.Get("/bands/:bandid/songs/", song.RouteGetSongsForBand)
	path := fmt.Sprintf("/bands/%v/songs/", "NotAnID")

	req, _ := http.NewRequest(http.MethodGet, path, nil)
	req.SetBasicAuth(s.ATUserEmail, s.ATUserPassword)

	res, err := app.Test(req, s.ATTimeout)
	s.Nil(err)
	s.Equal(res.StatusCode, 500)
}

func parseResponseBodyToArray(respBody []byte) []map[string]interface{} {
	var dat []map[string]interface{}
	if err := json.Unmarshal(respBody, &dat); err != nil {
		panic(err)
	}
	return dat

}

func parseResponseBodyToObject(respBody []byte) map[string]interface{} {
	var dat map[string]interface{}
	if err := json.Unmarshal(respBody, &dat); err != nil {
		panic(err)
	}
	return dat
}

func TestRunApp(t *testing.T) {
	suite.Run(t, new(TestAppSuite))
}
