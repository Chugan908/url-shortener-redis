package routes

import (
	"github.com/Chugan908/url-shortener-redis/database"
	"github.com/gofiber/fiber/v2"
)

// resolve.go is responsible for redirecting the user to the real url after pressing on the shortened one

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(url).Result()
	if value == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Short Url not found in the database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to database",
			"err":   err,
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr("counter")

	return c.Redirect(value, 301)
}
