package instrument

import (
	"app/config"
	"app/errors"
	"app/structs"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// RouteCreateInstrument is used to create an instrument for a given band.
func RouteCreateInstrument(c *fiber.Ctx) error {
	var instrument = new(structs.Instrument)
	if err := c.BodyParser(instrument); err != nil {
		return err
	}

	var err = CreateInstrument(instrument)

	if err != nil {
		return err
	}

	return c.SendString(config.Messages.InstrumentCreatedMessage)
}

// RouteDeleteInstrument is used to create an instrument for a given band.
func RouteDeleteInstrument(c *fiber.Ctx) error {
	instrumentID, err := strconv.Atoi(c.Params("instrumentid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	err = DeleteInstrument(instrumentID)
	if err != nil {
		return err
	}

	return c.SendString(config.Messages.InstrumentDeletedMessage)
}

// RouteGetInstrumentsForBand is used to collect all instruments for a given band.
func RouteGetInstrumentsForBand(c *fiber.Ctx) error {
	bandID, err := strconv.Atoi(c.Params("bandid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	instruments, err := GetInstrumentsForBand(bandID)
	if err != nil {
		return err
	}

	return c.JSON(instruments)
}
