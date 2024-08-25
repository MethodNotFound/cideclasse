package controllers

import (
  "github.com/gofiber/fiber/v2"
)

func JsonParseError(c *fiber.Ctx) error {
  return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
    "error": "Cannot parse JSON",
  })
}

func ValidateError(c *fiber.Ctx, err error) error {
  return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
    "error": err.Error(),
  })
}
