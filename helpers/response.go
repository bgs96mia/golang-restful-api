package helpers

import "github.com/gofiber/fiber/v2"

func ResponseWithError(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(
		&fiber.Map{"message": message},
	)
}

func ResponseWithSuccess(ctx *fiber.Ctx, statusCode int, message string, data interface{}) error {
	if data != nil {
		return ctx.Status(statusCode).JSON(
			&fiber.Map{"message": message, "data": data},
		)
	} else {
		return ctx.Status(statusCode).JSON(
			&fiber.Map{"message": message},
		)
	}
}
