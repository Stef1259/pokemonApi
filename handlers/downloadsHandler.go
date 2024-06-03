package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func DownloadFile(ctx *fiber.Ctx) error {
	return ctx.Download("./public/training_Manual.pdf")
}
