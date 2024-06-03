package handlers

import (
	"errors"
	"math/rand"
	"pokemonApi/database"
	"pokemonApi/model"

	"github.com/gofiber/fiber/v2"
)

func CreateTrainedPokemon(userID string) error {
	var count int
	if err := database.Database.Db.Model(&model.Pokemon{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return errors.New("no Pokemon species available")
	}

	randomNumber := rand.Intn(count)

	var randomPokemon model.Pokemon
	if err := database.Database.Db.Model(&model.Pokemon{}).Offset(randomNumber).First(&randomPokemon).Error; err != nil {
		return err
	}

	capturedPokemon := model.TrainerPokemon{
		Level:     1,
		UserID:    userID,
		PokemonID: randomNumber,
	}

	if err := database.TrainerPokemonDatabase.Db.Create(&capturedPokemon).Error; err != nil {
		return err
	}

	return nil
}

func GetAllcapturedPokemon(ctx *fiber.Ctx) error {
	var capturedPokemon []model.TrainerPokemon
	database.TrainerPokemonDatabase.Db.Find(&capturedPokemon)
	return ctx.JSON(capturedPokemon)
}

func GetcapturedPokemon(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var capturedPokemon model.TrainerPokemon
	database.TrainerPokemonDatabase.Db.First(&capturedPokemon, id)
	return ctx.JSON(capturedPokemon)
}

func UpdatecapturedPokemon(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	var capturedPokemon model.TrainerPokemon
	database.TrainerPokemonDatabase.Db.First(&capturedPokemon, id)
	if capturedPokemon.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "pokemons not found not found",
		})

	}
	updatedPokemon := new(model.TrainerPokemon)
	if err := ctx.BodyParser(updatedPokemon); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	database.TrainerPokemonDatabase.Db.Model(&capturedPokemon).Updates(updatedPokemon)
	return ctx.JSON(updatedPokemon)
}

func DeleteCpturedPokemon(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var trainer model.TrainerPokemon
	database.TrainerPokemonDatabase.Db.First(&trainer, id)
	if trainer.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "Pokemon not found",
		})
	}
	database.TrainerPokemonDatabase.Db.Delete(&trainer)
	return ctx.SendString("Pokemon  successfully released")
}
