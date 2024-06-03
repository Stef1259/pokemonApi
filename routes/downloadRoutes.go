package routes

import (
	"pokemonApi/handlers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RegisterDownloadRoutes(app *fiber.App) {
	app.Get("/download/tainingManual", handlers.AuthRequired(), limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Hour,
	}), handlers.DownloadFile)
}
