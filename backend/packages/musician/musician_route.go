package musician

import (
	"app/config"
	"app/errors"
	"app/structs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// RouteGetMusician returns a musician by ist email
func RouteGetMusician(c *fiber.Ctx) error {
	musicianEmail := c.Params("email")

	musician, err := GetMusicianByEmail(musicianEmail)
	if err != nil {
		return err
	}
	return c.JSON(musician)
}

// RouteCreateMusician is used to create a Musician (User) for the application.
func RouteCreateMusician(c *fiber.Ctx) error {
	var musician = new(structs.Musician)
	if err := c.BodyParser(musician); err != nil {
		return err
	}

	var err = CreateMusician(musician)

	if err != nil {
		return err
	}

	return c.SendString(config.Messages.MusicianCreatedMessage)
}

// RouteUpdateMusician is used to change Musician details.
func RouteUpdateMusician(c *fiber.Ctx) error {
	musicianID, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	var musician = new(structs.Musician)
	c.BodyParser(musician)

	err = UpdateMusician(*musician, musicianID)

	if err != nil {
		return err
	}

	return c.SendString(config.Messages.MusicianUpdatedMessage)
}
