package main

import (
  "ncdclssrm/models"
  "ncdclssrm/database"

  "github.com/go-playground/validator/v10"
  "github.com/gofiber/fiber/v2"
  "golang.org/x/crypto/bcrypt"
)

func main() {
  db := database.Connection
  app := fiber.New()

  app.Get("/logins", func(c *fiber.Ctx) error {
    var logins []models.Login
    db.Find(&logins)

    return c.JSON(fiber.Map{
      "data": logins,
    })
  })

  app.Post("/logins", func(c *fiber.Ctx) error {
    type request struct {
      Identificator  string `json:"identificator" validate:"required"`
      Password  string `json:"password" validate:"required"`
      Type  string `validate:"required"`
    }

    var validate = validator.New()

    req := new(request)
    req.Type = "Professor"

    if err := c.BodyParser(req); err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "error": "Cannot parse JSON",
      })
    }

    if err := validate.Struct(req); err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "error": err.Error(),
      })
    }

    // TODO add salt
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

    db.Create(&models.Login{Identificator: req.Identificator, PasswordHash: string(hashedPassword), Type: req.Type})

    return c.JSON(fiber.Map{
      "data": req,
    })
  })

  app.Listen(":3000")
}

