package initialize

import (
	"github.com/Chugan908/url-shortener-redis/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func NewApp() *fiber.App {
	return fiber.New()
}
