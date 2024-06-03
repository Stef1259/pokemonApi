package routes

import (
	"pokemonApi/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPokemonRouter(app *fiber.App) {
	app.Post("/pokemon", handlers.CreatePokemon)
	app.Get("/pokemon", handlers.GetAllPokemon)
	app.Get("/pokemon/:id", handlers.GetPokemon)
	app.Put("/pokemon/:id", handlers.UpdatePokemon)
	app.Delete("/pokemon/:id", handlers.DeletePokemon)
}
