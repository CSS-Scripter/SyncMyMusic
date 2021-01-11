package bandmember

import (
	"app/config"
	"app/errors"
	"app/structs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// RouteGetBandmembers returns a list of bandmembers
func RouteGetBandmembers(c *fiber.Ctx) error {
	bandID, err := strconv.Atoi(c.Params("bandid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	bandmembers, err := GetBandmembers(bandID)
	if err != nil {
		return err
	}

	return c.JSON(bandmembers)
}

// RouteUpdateBandmember updates a single bandmember
func RouteUpdateBandmember(c *fiber.Ctx) error {
	BandID, err := strconv.Atoi(c.Params("bandid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	MusicianID, err := strconv.Atoi(c.Params("musicianid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	bandmember := new(structs.BandMember)

	c.BodyParser(bandmember)
	bandmember.BandID = BandID
	bandmember.MusicianID = MusicianID

	err = UpdateBandmember(*bandmember)

	return err
}

// RouteDeleteBandmember deletes a single bandmember
func RouteDeleteBandmember(c *fiber.Ctx) error {
	BandID, err := strconv.Atoi(c.Params("bandid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	MusicianID, err := strconv.Atoi(c.Params("musicianid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	bandmember := new(structs.BandMember)

	c.BodyParser(bandmember)
	bandmember.BandID = BandID
	bandmember.MusicianID = MusicianID

	err = DeleteBandmember(*bandmember)

	if err != nil {
		return err
	}

	return c.SendString(config.Messages.BandmemberDeletedMessage)
}
