package main

import (
	"pokemonApi/database"
	"pokemonApi/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	database.ConnectTrainerDb()
	database.ConnectTrainerPokemonDb()

	app := fiber.New()

	app.Static("/", "pokemonApi/public")

	routes.RegisterPokemonRouter(app)
	routes.RegisterTrainerRouter(app)
	routes.RegisterCapturedPokemonRouter(app)
	routes.RegisterDownloadRoutes(app)
	routes.RegisterAuthRoutes(app)

	app.Listen(":3000")
}
