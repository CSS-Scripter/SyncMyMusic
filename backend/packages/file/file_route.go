package file

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

// RouteGetFile returns the requested file
func RouteGetFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	bandID := c.Params("bandId")
	songID := c.Params("songId")
	pdfLocation := fmt.Sprintf("./files/%v/%v/%v/%v", os.Getenv("APP_ENV"), bandID, songID, filename)
	return c.SendFile(pdfLocation)
}
