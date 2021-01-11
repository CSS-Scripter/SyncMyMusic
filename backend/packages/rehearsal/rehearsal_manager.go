package rehearsal

import (
	"app/errors"
	"app/packages/band"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var rehearsals map[int]Rehearsal

func init() {
	rehearsals = map[int]Rehearsal{}
}

func startRehearsal(bandID int, ownerID int) *Rehearsal {
	rehearsal := NewRehearsal(bandID, ownerID)
	go rehearsal.pool.Start()
	rehearsals[bandID] = *rehearsal
	return rehearsal
}

// JoinRehearsal calls Startrehearsal if rehearsal does not exist, then joins rehearsal
func JoinRehearsal(c *websocket.Conn) error {

	bandID, err := strconv.Atoi(c.Params("bandID"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}

	rehearsal, exists := getRehearsal(bandID)

	if !exists {
		rehearsal = *startRehearsal(bandID, userID)
	}

	client := &Client{
		ID:   userID,
		Conn: c,
		Pool: rehearsal.pool,
	}

	rehearsal.pool.Register <- client
	client.Read()

	return nil
}

func getRehearsal(id int) (Rehearsal, bool) {
	rehearsal, ok := rehearsals[id]
	return rehearsal, ok
}

// DeleteRehearsal Remove rehearsal from list by id
func DeleteRehearsal(id int) {
	delete(rehearsals, id)
}

// RouteGetRehearsals returns a list of current rehearsals, based on the requesting users id
func RouteGetRehearsals(c *fiber.Ctx) error {
	userID := c.Params("userid")
	userIDAsInt, err := strconv.Atoi(userID)
	if err != nil {
		return new(errors.IDIsNotANumberError)
	}
	bands, err := band.GetBandsUserIsIn(userIDAsInt)
	if err != nil {
		return err
	}

	result := []int{}
	for _, band := range bands {
		if _, ok := rehearsals[band.ID]; ok {
			result = append(result, band.ID)
		}
	}

	return c.JSON(result)
}
