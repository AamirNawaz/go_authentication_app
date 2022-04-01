package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func MUsers(c *fiber.Ctx) error {
	return c.SendString("Mongo db users")
}
