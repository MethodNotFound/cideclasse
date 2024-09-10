package main

import (
  "cideclasse/controllers"
  "cideclasse/database"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/template/html/v2"

  jwtware "github.com/gofiber/contrib/jwt"
  // "github.com/golang-jwt/jwt/v5"
)

func main() {
  database.Setup()

  engine := html.New("backend/views", ".html")

  app := fiber.New(fiber.Config{
    Views: engine,
    ViewsLayout: "layouts/main",
    PassLocalsToViews: true,
  })

  controllers.DefineSessionsEndPoints(app)

  app.Static("/", "./public")

  app.Get("/front/home", func(c *fiber.Ctx) error {
    return c.Render("dashboard", fiber.Map{
      "Title": "Hdello, World!",
    })
  })

  app.Use(jwtware.New(jwtware.Config{
    SigningKey: jwtware.SigningKey{Key: []byte("secret")},
  }))

  controllers.DefineAdminStudentsEndPoints(app)
  controllers.DefineAdminClassesEndPoints(app)
  controllers.DefineStudentsEndPoints(app)

  app.Listen(":3000")
}

