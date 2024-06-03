package routes

import (
	"pokemonApi/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterTrainerRouter(app *fiber.App) {
	app.Post("/trainer", handlers.CreateTrainer)
	app.Get("/trainer", handlers.GetAllTrainer)
	app.Get("/trainer/:id", handlers.GetTrainer)
	app.Put("/trainer/:id", handlers.UpdateTrainer)
	app.Delete("/trainer/:id", handlers.DeleteTrainer)
	app.Get("/trainer/:id/capture", handlers.CapturePokemon)
}
