package main

import (
	"log"
	"os"

	"github.com/Chugan908/url-shortener-redis/initialize"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := initialize.NewApp()

	app.Use(logger.New())

	initialize.SetupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
