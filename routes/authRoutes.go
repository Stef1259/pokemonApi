package routes

import (
	"pokemonApi/handlers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RegisterAuthRoutes(app *fiber.App) {
	app.Get("/login", limiter.New(limiter.Config{
		Max:        4,
		Expiration: 1 * time.Hour,
	}), handlers.Login)
}
