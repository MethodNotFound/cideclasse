package controllers

import (
  "cideclasse/database"
  "cideclasse/models"

  // "github.com/go-playground/validator/v10"
  "github.com/gofiber/fiber/v2"
  "github.com/golang-jwt/jwt/v5"
)

func DefineStudentsEndPoints(app *fiber.App) {
  db := database.Connection

  app.Get("/current_student", func(c *fiber.Ctx) error {
    sessions := c.Locals("user").(*jwt.Token)
    claims := sessions.Claims.(jwt.MapClaims)
    session_id := claims["id"].(string)

    var session models.Session

    result := db.Where("id = ?", session_id).First(&session)
    if result.Error != nil {
      return c.JSON(fiber.Map{
        "error": result.Error,
      })
    } else {
      return c.JSON(fiber.Map{
        "data": session,
      })
    }
  })
}
