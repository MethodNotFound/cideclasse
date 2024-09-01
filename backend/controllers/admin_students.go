package controllers

import (
  "cideclasse/database"
  "cideclasse/models"

  "github.com/go-playground/validator/v10"
  "github.com/gofiber/fiber/v2"
)

func DefineAdminStudentsEndPoints(app *fiber.App) {
  db := database.Connection


  app.Patch("/admin/students/:identifier/reset", func(c *fiber.Ctx) error {
    err := RequireAdmin(c)
    if err != nil {
      return err
    }

    student, err := models.FindStudent(db, c.Params("identifier"))
    if err != nil {
      return c.JSON(fiber.Map{
        "error": err,
      })
    }
    student.AskNewPassword = true
    db.Save(&student)
    return c.JSON(fiber.Map{
      "data": student,
    })
  })

  app.Post("/admin/students", func(c *fiber.Ctx) error {
    err := RequireAdmin(c)
    if err != nil {
      return err
    }

    type request struct {
      Identifier  string `json:"identifier" validate:"required"`
      Name  string `json:"name" validate:"required"`
    }

    req := new(request)
    if err := c.BodyParser(req); err != nil {
      return JsonParseError(c)
    }
    if err := validator.New().Struct(req); err != nil {
      return ValidateError(c, err)
    }

    student, err := models.CreateStudent(db, req.Name, req.Identifier)
    if err != nil {
      return c.JSON(fiber.Map{
        "error": err,
      })
    }

    return c.JSON(fiber.Map{
      "success": "student created",
      "data": student,
    })
  })
}
