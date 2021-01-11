package main

import (
	"app/packages/bandmember"
	"app/packages/file"
	"app/packages/rehearsal"
	"os"
	"strings"

	"app/packages/band"
	"app/packages/instrument"
	"app/packages/musician"
	"app/packages/song"
	"app/packages/songinstrument"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	setupMiddleware(app)
	setupRoutes(app)

	app.Listen(":3000")
}

func setupMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(logger.New())
	setupAuth(app)
}

func setupAuth(app *fiber.App) {
	app.Use(basicauth.New(basicauth.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/register" || strings.Contains(c.Path(), "/login") || strings.Contains(c.Path(), "/ws") || c.Path() == "/"
		},
		Authorizer: func(email, password string) bool {
			return musician.AuthenticatUser(email, password)
		},
		ContextUsername: "email",
	}))

}

func setupRoutes(app *fiber.App) {
	app.Get("/env", func(c *fiber.Ctx) error {
		return c.SendString(os.Getenv("APP_ENV"))
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:bandID", websocket.New(func(c *websocket.Conn) {
		rehearsal.JoinRehearsal(c)
	}))

	app.Static("/coverage", "./coverage.html")

	app.Get("users/:userid/rehearsals", rehearsal.RouteGetRehearsals)
	app.Post("/register", musician.RouteCreateMusician)
	app.Get("/login/:email", musician.RouteGetMusician)

	app.Get("/users/:userid/bands", band.RouteGetBandsUserIsIn)
	app.Post("/users/:userid/bands", band.RouteCreateBand)
	app.Delete("/bands/:bandid", band.RouteDeleteBand)
	app.Get("/bands/:bandid", band.RouteGetSingleBand)
	app.Post("/bands/:bandid/invite", band.RouteAddUserToBand)

	app.Post("/bands/:bandid/request", band.RouteJoinRequest)
	app.Get("/bands/:bandid/request", band.RouteGetJoinRequestsForBand)
	app.Put("/bands/:bandid/request/:userid/accept", band.RouteAcceptJoinRequest)
	app.Put("/bands/:bandid/request/:userid/deny", band.RouteDenyJoinRequest)

	app.Put("/bands/:bandid/songs/:songid", song.RouteUpdateSong)
	app.Get("/bands/:bandid/songs", song.RouteGetSongsForBand)
	app.Post("/bands/:bandid/songs", song.RouteCreateSongForBand)
	app.Delete("/songs/:songid", song.RouteDeleteSong)

	app.Get("/bands/:bandid/instruments", instrument.RouteGetInstrumentsForBand)
	app.Post("/bands/:bandid/instruments", instrument.RouteCreateInstrument)
	app.Delete("/bands/:bandid/instruments/:instrumentid", instrument.RouteDeleteInstrument)
	app.Get("/songs/:songid/instruments", song.RouteGetInstrumentsForSong)

	app.Post("/songs/:songid/instruments/:instrumentid", songinstrument.RouteUploadFile)
	app.Put("/songs/:songid/instruments/:instrumentid", songinstrument.RouteUpdateFile)

	app.Get("/files/:bandId/:songId/:filename", file.RouteGetFile)

	app.Get("/bands/:bandid/members", bandmember.RouteGetBandmembers)
	app.Put("/bands/:bandid/members/:musicianid", bandmember.RouteUpdateBandmember)
	app.Delete("/bands/:bandid/members/:musicianid", bandmember.RouteDeleteBandmember)
	app.Get("/members/:email", musician.RouteGetMusician)
	app.Put("/users/:userid", musician.RouteUpdateMusician)
}
