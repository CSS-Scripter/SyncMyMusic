package songinstrument

import (
	"app/structs"
	"fmt"
	"os"

	song "app/packages/song"
	dao "app/packages/songinstrument/dao"

	"github.com/gofiber/fiber/v2"
)

var songGetBandIDFromSong = song.GetBandIDFromSong
var daoCreateSongInstrument = dao.CreateSongInstrument

// CreateSongInstrument creates an entry for
func CreateSongInstrument(songID int, instrumentID int) (int, error) {
	bandID, err := songGetBandIDFromSong(songID)
	if err != nil {
		return 0, err
	}
	path := fmt.Sprintf("./files/%v/%v/%v.pdf", bandID, songID, instrumentID)
	songInstrument := &structs.SongInstrument{
		SongID:       songID,
		InstrumentID: instrumentID,
		PDFLocation:  path,
	}
	return bandID, daoCreateSongInstrument(songInstrument)
}

func saveFile(c *fiber.Ctx, location string, filename string) error {
	file, err := c.FormFile("document")
	if err != nil {
		return err
	}
	fileType := file.Header["Content-Type"]
	if len(fileType) > 0 && fileType[0] == "application/pdf" {
		os.MkdirAll(location, os.ModePerm)
		err := c.SaveFile(file, location+filename)
		if err != nil {
			return err
		}
		return c.SendString("Success")
	}
	return c.Status(422).SendString("Received file was not a PDF")
}
