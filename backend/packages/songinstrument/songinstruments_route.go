package songinstrument

import (
	"fmt"
	"os"
	"strconv"

	"app/errors"
	"app/packages/song"

	"github.com/gofiber/fiber/v2"
)

// RouteUploadFile uploads a file. Path POST
func RouteUploadFile(c *fiber.Ctx) error {
	songID, err := strconv.Atoi(c.Params("songid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}
	instrumentID, err := strconv.Atoi(c.Params("instrumentid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}
	bandID, err := CreateSongInstrument(songID, instrumentID)
	if err != nil {
		return err
	}
	location := fmt.Sprintf("./files/%v/%v/%v/", os.Getenv("APP_ENV"), bandID, songID)
	filename := fmt.Sprintf("%v.pdf", instrumentID)
	return saveFile(c, location, filename)
}

// RouteUpdateFile uploads a file. Path POST
func RouteUpdateFile(c *fiber.Ctx) error {
	songID, err := strconv.Atoi(c.Params("songid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}
	instrumentID, err := strconv.Atoi(c.Params("instrumentid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}
	bandID, err := song.GetBandIDFromSong(songID)
	if err != nil {
		return err
	}
	pdfLocation := fmt.Sprintf("./files/%v/%v/%v/", os.Getenv("APP_ENV"), bandID, songID)
	filename := fmt.Sprintf("%v.pdf", instrumentID)
	return saveFile(c, pdfLocation, filename)
}
