package song

import (
	"fmt"
	"strconv"

	"app/config"
	"app/errors"
	"app/structs"

	"github.com/gofiber/fiber/v2"
)

// RouteCreateSongForBand is the function called when an bandmember wants to create a song. POST /bands/:bandid/songs
func RouteCreateSongForBand(c *fiber.Ctx) error {
	bandID, err := strconv.Atoi(c.Params("bandid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	var song = new(structs.Song)
	if err := c.BodyParser(song); err != nil {
		return err
	}

	ID, err := CreateSongForBand(bandID, song)
	if err != nil {
		return err
	}

	return c.SendString(fmt.Sprint(ID))
}

// RouteGetSongsForBand returns all songs found for bandid. GET: /bands/:bandid/songs
func RouteGetSongsForBand(c *fiber.Ctx) error {
	bandID, err := strconv.Atoi(c.Params("bandid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	songs, err := GetSongsForBand(bandID)
	if err != nil {
		return err
	}

	return c.JSON(songs)
}

// RouteDeleteSong is called when a song needs to be deleted. DELETE: /songs/:songid
func RouteDeleteSong(c *fiber.Ctx) error {
	songID, err := strconv.Atoi(c.Params("songid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	err = DeleteSong(songID)
	if err != nil {
		return err
	}

	return c.SendString(config.Messages.SongDeletedMessage)
}

// RouteUpdateSong updates an song. Called on PUT /bands/:bandid/songs/:songid
func RouteUpdateSong(c *fiber.Ctx) error {
	songID, err := strconv.Atoi(c.Params("songid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	var song = new(structs.Song)
	if err := c.BodyParser(song); err != nil {
		return err
	}
	song.ID = songID

	return UpdateSong(song)
}

// RouteGetInstrumentsForSong is called on path GET /songs/:songid/instruments
func RouteGetInstrumentsForSong(c *fiber.Ctx) error {
	songID, err := strconv.Atoi(c.Params("songid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	instruments, err := GetInstrumentsForSong(songID)
	if err != nil {
		return err
	}

	return c.JSON(instruments)
}
