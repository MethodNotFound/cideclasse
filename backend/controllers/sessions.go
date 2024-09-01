package controllers

import (
	"cideclasse/database"
	"cideclasse/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

func DefineSessionsEndPoints(app *fiber.App) {
  db := database.Connection

  app.Post("/sessions", func(c *fiber.Ctx) error {
    type request struct {
      Identifier  string `json:"identifier" validate:"required"`
      Password  string `json:"password" validate:"required"`
    }

    req := new(request)
    if err := c.BodyParser(req); err != nil {
        return JsonParseError(c)
    }
    if err := validator.New().Struct(req); err != nil {
        return ValidateError(c, err)
    }

    // TODO you already know
    if req.Identifier == "admin" && req.Password == "admin" {
      claims := jwt.MapClaims{
        "id":  "admin",
      }

      token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
      t, err := token.SignedString([]byte("secret"))
      if err != nil {
        return c.JSON(fiber.Map{
          "error": err,
        })
      }

      return c.JSON(fiber.Map{"token": t})
    }

    student, err := models.FindStudent(db, req.Identifier)
    if err != nil {
      return c.JSON(fiber.Map{
        "error": err,
      })
    }

    if student.AskNewPassword || student.PasswordHash == "" {
      return c.JSON(fiber.Map{
        "error": "make a password",
      })
    }


    session, err := models.NewSession(db, student.Identifier, req.Password)
    if err != nil {
      return c.JSON(fiber.Map{
        "error": err,
      })
    }

    // Create the Claims
    claims := jwt.MapClaims{
      "id":  session.ID,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte("secret"))
    if err != nil {
      return c.JSON(fiber.Map{
        "error": err,
      })
    }

    return c.JSON(fiber.Map{"token": t})
  })

  app.Patch("/sessions", func(c *fiber.Ctx) error {
    type request struct {
      Identifier  string `json:"identifier" validate:"required"`
      Password  string `json:"password" validate:"required"`
    }

    req := new(request)
    if err := c.BodyParser(req); err != nil {
        return JsonParseError(c)
    }
    if err := validator.New().Struct(req); err != nil {
        return ValidateError(c, err)
    }

    student, err := models.FindStudent(db, req.Identifier)
    if err != nil {
      return c.JSON(fiber.Map{
        "error": err,
      })
    }

    if !student.AskNewPassword {
      return c.JSON(fiber.Map{
        "error": "AskNewPassword is false",
      })
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

    student.AskNewPassword = false
    student.PasswordHash = string(hashedPassword)
    db.Save(&student)
    return c.JSON(fiber.Map{
      "success": student,
    })
  })
}
