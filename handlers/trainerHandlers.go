package handlers

import (
	"github.com/gofiber/fiber/v2"
	"pokemonApi/database"
	"pokemonApi/model"
)

func CreateTrainer(ctx *fiber.Ctx) error {

	// Create a new Pokemon
	trainer := new(model.Trainer)

	if err := ctx.BodyParser(trainer); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	database.TrainerDatabase.Db.Create(&trainer)
	return ctx.JSON(trainer)

}

func GetAllTrainer(ctx *fiber.Ctx) error {
	var trainer []model.Trainer
	database.TrainerDatabase.Db.Find(&trainer)
	return ctx.JSON(trainer)
}

func GetTrainer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var trainer model.Trainer
	database.TrainerDatabase.Db.First(&trainer, id)
	return ctx.JSON(trainer)
}

func UpdateTrainer(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	var trainer model.Trainer
	database.TrainerDatabase.Db.First(&trainer, id)
	if trainer.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "Trainer not found",
		})

	}
	updatedTrainer := new(model.Trainer)
	if err := ctx.BodyParser(updatedTrainer); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	database.TrainerDatabase.Db.Model(&trainer).Updates(updatedTrainer)
	return ctx.JSON(updatedTrainer)
}

func DeleteTrainer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var trainer model.Trainer
	database.TrainerDatabase.Db.First(&trainer, id)
	if trainer.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "trainer not found",
		})
	}
	database.TrainerDatabase.Db.Delete(&trainer)
	return ctx.SendString("Trainer successfully deleted")
}

func GetAllTrainerPokemon(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var trainerPokemon []model.TrainerPokemon
	database.TrainerPokemonDatabase.Db.Where("user_id = ?", id).Find(&trainerPokemon)
	return ctx.JSON(trainerPokemon)
}

func GetTrainerPokemon(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var trainerPokemon model.TrainerPokemon
	database.TrainerPokemonDatabase.Db.First(&trainerPokemon, id)
	return ctx.JSON(trainerPokemon)
}

func CapturePokemon(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := CreateTrainedPokemon(id)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "pokemon escaped",
		})
	}

	return ctx.SendString("Pokemon captured")
}
