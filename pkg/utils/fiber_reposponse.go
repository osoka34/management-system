package utils

import "github.com/gofiber/fiber/v3"

func ValidationErr(err error) fiber.Map {
	return fiber.Map{
		"error":   "Validation failed",
		"details": err.Error(),
	}
}

func BadRequestErr(err error) fiber.Map {
	return fiber.Map{
		"error":   "Invalid request body",
		"details": err.Error(),
	}
}

func InternalErr(err error) fiber.Map {
	return fiber.Map{
		"error":   "Internal server error",
		"details": err.Error(),
	}
}
