package routes

import (
	"pokemonApi/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterCapturedPokemonRouter(app *fiber.App) {
	app.Get("/captured", handlers.GetAllPokemon)
	app.Get("/captured/:id", handlers.GetcapturedPokemon)
	app.Put("/captured/:id", handlers.UpdatecapturedPokemon)
	app.Delete("/captured/:id", handlers.DeleteCpturedPokemon)
}
