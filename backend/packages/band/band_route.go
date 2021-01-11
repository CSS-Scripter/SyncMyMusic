package band

import (
	"fmt"
	"strconv"

	"app/errors"
	"app/structs"

	"app/config"

	"github.com/gofiber/fiber/v2"
)

// RouteGetSingleBand is the function called for endpoint GET /bands/:bandid
func RouteGetSingleBand(c *fiber.Ctx) error {
	bandID := c.Params("bandid")
	bandIDAsInt, err := strconv.Atoi(bandID)
	if err != nil {
		return err
	}
	band, err := GetBand(bandIDAsInt)
	if err != nil {
		return err
	}
	return c.JSON(band)
}

// RouteAddUserToBand adds a new Musician to a band
func RouteAddUserToBand(c *fiber.Ctx) error {
	bandID := c.Params("bandid")
	bandIDAsInt, err := strconv.Atoi(bandID)
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	var musician = new(structs.Musician)
	if err := c.BodyParser(musician); err != nil {
		return err
	}
	err = AddUserToBand(bandIDAsInt, musician)
	if err != nil {
		return err
	}
	return c.Status(201).SendString("ok")
}

// RouteGetBandsUserIsIn is the function called for endpoint GET /users/:userid/bands
func RouteGetBandsUserIsIn(c *fiber.Ctx) error {
	userID := c.Params("userid")
	userIDAsInt, err := strconv.Atoi(userID)
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}
	bands, err := GetBandsUserIsIn(userIDAsInt)
	if err != nil {
		return err
	}
	return c.JSON(bands)
}

// RouteCreateBand is the function called when an user wants to create a band. Endpoint POST: /users/:userid/bands
func RouteCreateBand(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}
	var band = new(structs.Band)
	if err := c.BodyParser(band); err != nil {
		return err
	}
	bandID, err := CreateBandForUser(userID, band)
	if err != nil {
		return err
	}
	return c.JSON(bandID)
}

// RouteDeleteBand is called when the user wants to delete a band. Endpoint DELETE: /bands/:bandid
func RouteDeleteBand(c *fiber.Ctx) error {
	bandID, err := strconv.Atoi(c.Params("bandid"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	email := fmt.Sprint(c.Locals("email"))

	status, err := DeleteBand(bandID, email)
	if err != nil {
		return err
	}

	return c.Status(status).SendString(config.Messages.BandDeletedMessage)
}

// RouteJoinRequest is called when the user wants to join a band. Enpoint POST: /bands/:bandid/request
func RouteJoinRequest(c *fiber.Ctx) error {
	bandID, err := strconv.Atoi(c.Params("bandid"))
	if err != nil {
		return new(errors.InvalidIDError)
	}
	email := fmt.Sprintf("%+v", c.Locals("email"))
	err = JoinRequest(email, bandID)
	if err != nil {
		return err
	}
	return c.SendString("Request send succesfully")
}

// RouteGetJoinRequestsForBand gets all join requests for a certain band
// Requires authenticator to be the bandowner.
func RouteGetJoinRequestsForBand(c *fiber.Ctx) error {
	bandID, err := strconv.Atoi(c.Params("bandid"))
	if err != nil {
		return new(errors.InvalidIDError)
	}
	email := fmt.Sprintf("%+v", c.Locals("email"))
	joinRequests, statusCode, err := GetJoinRequestsForBand(email, bandID)
	if err != nil {
		return err
	}
	return c.Status(statusCode).JSON(joinRequests)
}

// RouteAcceptJoinRequest is called when a join request is accepted
func RouteAcceptJoinRequest(c *fiber.Ctx) error {
	bandID, err := strconv.Atoi(c.Params("bandid"))
	if err != nil {
		return new(errors.InvalidIDError)
	}
	userID, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return new(errors.InvalidIDError)
	}
	email := fmt.Sprintf("%+v", c.Locals("email"))
	status, err := AcceptJoinRequest(email, bandID, userID)
	if err != nil {
		return err
	}
	return c.Status(status).SendString("")
}

// RouteDenyJoinRequest is called when a join request is denied
func RouteDenyJoinRequest(c *fiber.Ctx) error {
	bandID, err := strconv.Atoi(c.Params("bandid"))
	if err != nil {
		return new(errors.InvalidIDError)
	}
	userID, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return new(errors.InvalidIDError)
	}
	email := fmt.Sprintf("%+v", c.Locals("email"))
	status, err := DenyJoinRequest(email, bandID, userID)
	if err != nil {
		return err
	}
	return c.Status(status).SendString("")
}
