package handlers

import (
	"pokemonApi/database"
	"pokemonApi/model"

	"github.com/gofiber/fiber/v2"
)

func CreatePokemon(ctx *fiber.Ctx) error {

	// Create a new Pokemon
	pokemon := new(model.Pokemon)

	if err := ctx.BodyParser(pokemon); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	database.Database.Db.Create(&pokemon)
	return ctx.JSON(pokemon)

}

func GetAllPokemon(ctx *fiber.Ctx) error {
	var pokemon []model.Pokemon
	database.Database.Db.Find(&pokemon)
	return ctx.JSON(pokemon)
}

func GetPokemon(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var pokemon model.Pokemon
	database.Database.Db.First(&pokemon, id)
	return ctx.JSON(pokemon)
}

func UpdatePokemon(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	var pokemon model.Pokemon
	database.Database.Db.First(&pokemon, id)
	if pokemon.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "Pokemon not found",
		})

	}
	updatedPokemon := new(model.Pokemon)
	if err := ctx.BodyParser(updatedPokemon); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	database.Database.Db.Model(&pokemon).Updates(updatedPokemon)
	if database.Database.Db.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": database.Database.Db.Error.Error(),
		})
	}
	return ctx.JSON(updatedPokemon)
}

func DeletePokemon(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var pokemon model.Pokemon
	database.Database.Db.First(&pokemon, id)
	if pokemon.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "Pokemon not found",
		})
	}
	database.Database.Db.Delete(&pokemon)
	return ctx.SendString("Pokemon successfully deleted")
}
