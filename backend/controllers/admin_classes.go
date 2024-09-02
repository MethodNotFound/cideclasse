package controllers

import (
  "cideclasse/database"
  "cideclasse/models"

  "github.com/go-playground/validator/v10"
  "github.com/gofiber/fiber/v2"
)

func DefineAdminClassesEndPoints(app *fiber.App) {
  db := database.Connection

  app.Get("/admin/classes/:name", func(c *fiber.Ctx) error {
    err := RequireAdmin(c)
    if err != nil {
      return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": err,
      })
    }

    class, err := models.FindClass(db, c.Params("name"))
    if err != nil {
      return c.JSON(fiber.Map{
        "error": err,
      })
    }
    return c.JSON(fiber.Map{
      "data": class,
    })
  })

  app.Get("/admin/classes", func(c *fiber.Ctx) error {
    err := RequireAdmin(c)
    if err != nil {
      return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": err,
      })
    }

    var classes []models.Class

    result := db.Find(&classes)
    if result.Error != nil {
      return c.JSON(fiber.Map{
        "error": result.Error,
      })
    }

    return c.JSON(fiber.Map{
      "data": classes,
    })
  })

  app.Post("/admin/classes", func(c *fiber.Ctx) error {
    err := RequireAdmin(c)
    if err != nil {
      return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": err,
      })
    }

    type request struct {
      Name  string `json:"name" validate:"required"`
    }

    req := new(request)
    if err := c.BodyParser(req); err != nil {
      return JsonParseError(c)
    }
    if err := validator.New().Struct(req); err != nil {
      return ValidateError(c, err)
    }

    class, err := models.CreateClass(db, req.Name)
    if err != nil {
      return c.JSON(fiber.Map{
        "error": err,
      })
    }

    return c.JSON(fiber.Map{
      "success": "class created",
      "data": class,
    })
  })
}
