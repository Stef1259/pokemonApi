package handlers

import (
	"pokemonApi/database"
	"pokemonApi/model"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
			return nil
		},
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	})
}

func Login(ctx *fiber.Ctx) error {

	trainerMail := ctx.Params("userEmail")

	if trainerMail != "" {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Email is required",
			"email":   trainerMail,
		})

	}

	var data map[string]string

	err := ctx.BodyParser(&data)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if data["password"] == "" {

		return ctx.Status(400).JSON(fiber.Map{
			"message": "Password is required",
		})

	}

	var trainer model.Trainer

	database.TrainerDatabase.Db.Where("email = ?", data["userEmail"]).First(&trainer)

	if trainer.Email == "" {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Trainer not found",
			"email":   trainerMail,
		})

	}

	if trainer.Password != data["password"] {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer": trainer.Email,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	trainerData := make(map[string]interface{})
	trainerData["token"] = tokenString

	return ctx.Status(200).JSON(fiber.Map{
		"token":   tokenString,
		"message": "Login successful",
	},
	)

}
