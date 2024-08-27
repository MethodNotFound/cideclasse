package controllers

import (
	"cideclasse/database"
	"cideclasse/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func DefineLoginEndPoints(app *fiber.App) {
  db := database.Connection

  app.Get("/logins", func(c *fiber.Ctx) error {
    var logins []models.Login
    db.Find(&logins)

    return c.JSON(fiber.Map{
      "data": logins,
    })
  })


  app.Post("/login_professor", func(c *fiber.Ctx) error {
    type request struct {
      Identificator  string `json:"identificator" validate:"required"`
      Password  string `json:"password" validate:"required"`
    }

    req := new(request)
    if err := c.BodyParser(req); err != nil {
        return JsonParseError(c)
    }
    if err := validator.New().Struct(req); err != nil {
        return ValidateError(c, err)
    }

    var login models.Login
    result := db.Where("identificator = ?", req.Identificator).First(&login)
    if result.Error != nil {
      if result.Error == gorm.ErrRecordNotFound {
        return c.JSON(fiber.Map{
          "error": "login not found",
        })
      }
      return c.JSON(fiber.Map{
        "error": result.Error,
      })
    }

    if err := bcrypt.CompareHashAndPassword([]byte(login.PasswordHash), []byte(req.Password)); err != nil {
      return c.JSON(fiber.Map{
        "error": "password incorrect",
        "pass_h": login.PasswordHash,
      })
    }
    return c.JSON(fiber.Map{
      "data": "logged in",
    })
  })

  app.Post("/logins", func(c *fiber.Ctx) error {
    type request struct {
      Identificator  string `json:"identificator" validate:"required,min=6"`
      Password  string `json:"password" validate:"required,min=6"`
      Type  string `validate:"required"`
    }
    req := new(request)
    req.Type = "Professor"

    if err := c.BodyParser(req); err != nil {
        return JsonParseError(c)
    }
    if err := validator.New().Struct(req); err != nil {
        return ValidateError(c, err)
    }

    var count int64
    db.Model(&models.Login{}).Where("identificator = ?", req.Identificator).Count(&count)
    if count > 0 {
      return c.JSON(fiber.Map{
        "error": "user already exists",
      })
    }

    // TODO add salt
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

    db.Create(&models.Login{Identificator: req.Identificator, PasswordHash: string(hashedPassword), Type: req.Type})

    return c.JSON(fiber.Map{
      "data": req,
    })
  })
}
