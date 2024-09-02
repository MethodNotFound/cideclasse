package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
)

func RequireAdmin(c *fiber.Ctx) error {
  session := c.Locals("user").(*jwt.Token)
  claims := session.Claims.(jwt.MapClaims)
  id := claims["id"].(string)

  if id != "admin" {
    return errors.New("not admin")
  }
  return nil
}


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
