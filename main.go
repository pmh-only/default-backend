package main

import (
	"log"
	"path"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	themeDir := getThemeDirectory()
	log.Printf("theme: %s loaded\n", themeDir)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,

		ViewsLayout: "layouts/default",
		Views:       handlebars.New(themeDir, ".hbs"),
	})

	app.Use(logger.New())
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
		})
	})

	app.Static("/", path.Join(themeDir, "assets"))
	app.All("/*", func(c *fiber.Ctx) error {
		statusCode, err := strconv.ParseInt(c.Get("X-Code", "404"), 10, 64)
		if err != nil {
			statusCode = 404
		}

		themeView := getThemeViewFile(themeDir, statusCode)
		return c.Render(themeView, fiber.Map{
			"status":  statusCode,
			"headers": c.GetReqHeaders(),
		})
	})

	log.Fatal(app.Listen(":3000"))
}
