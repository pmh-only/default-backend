package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(logger.New())
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
		})
	})

	app.All("/*", func(c *fiber.Ctx) error {
		statusCode, err := strconv.ParseInt(c.Get("X-Code", "400"), 10, 64)
		if err != nil {
			statusCode = 400
		}

		return c.Status(int(statusCode)).JSON(fiber.Map{
			"headers": c.GetReqHeaders(),
			"path":    c.Path(),
		})
	})

	log.Fatal(app.Listen(":3000"))
}
